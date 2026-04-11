package review

import (
	card_state_domain "card/internal/domain/card_state"
	review_domain "card/internal/domain/review"
	"math"
	"time"
)

type SchedulerInterface interface {
	ReviewCard(cardState *card_state_domain.CardState, rating review_domain.RatingType, reviewDate time.Time) (*review_domain.ReviewLog, error)
}

type StateFunc = func(cardState *card_state_domain.CardState, rating review_domain.RatingType, reviewDate time.Time)

type Scheduler struct {
	learningSteps            []time.Duration // Small-time intervals that schedule cards in the Learning state
	relearningSteps          []time.Duration // Small-time intervals that schedule cards in the Relearning state.
	minimalIntervalInDays    int64           // The minimum interval (in days) given to a review card after answering Again.
	maximumIntervalInDays    int64           // The maximum number of days a Review-state card can be scheduled into the future.
	startingEasy             float64         // The initial ease factor given to cards that have completed the learning steps and become a Review-state card.
	easyBonus                float64         // An extra multiplier that is applied to a review card's interval when you rate it Easy.
	intervalModifier         float64         // A factor used as a multiplier to determine future review interval lengths. It is used on Review-state cards and Relearning-state cards about to graduate the relearning steps.
	easyIntervalInDays       int64           // The number of days to wait before showing a card again, after the Easy button is used to immediately remove a card from learning.
	graduationIntervalInDays int64           // The number of days to wait before showing a card again, after the Good button is pressed on the final learning step.
	hardInterval             float64         // The multiplier applied to a review interval when answering Hard.
	newInterval              float64         // The multiplier applied to a review interval when answering Again.
	states                   map[card_state_domain.CardStateType]StateFunc
	fuzzy                    *fuzzy
}

func NewScheduler() *Scheduler {
	s := &Scheduler{
		learningSteps:            []time.Duration{time.Minute, 10 * time.Minute},
		relearningSteps:          []time.Duration{10 * time.Minute},
		minimalIntervalInDays:    1,
		maximumIntervalInDays:    36500, // 100 years
		startingEasy:             2.5,
		easyBonus:                1.3,
		intervalModifier:         1.0,
		easyIntervalInDays:       4,
		graduationIntervalInDays: 1,
		hardInterval:             1.2,
		newInterval:              0.0,
		fuzzy:                    newFuzzy(1, 36500),
	}

	s.states = map[card_state_domain.CardStateType]StateFunc{
		card_state_domain.LearningCardState:   s.stateLearningFunc,
		card_state_domain.ReviewCardState:     s.stateReviewFunc,
		card_state_domain.RelearningCardState: s.stateRelearningFunc,
	}

	return s
}

func (s *Scheduler) ReviewCard(cardState *card_state_domain.CardState, rating review_domain.RatingType, reviewDate time.Time) (*review_domain.ReviewLog, error) {
	reviewLog, err := review_domain.NewReviewLog(cardState.Id, rating, reviewDate)

	s.states[cardState.State](cardState, rating, reviewDate)

	return reviewLog, err
}

func (s *Scheduler) stateLearningFunc(cardState *card_state_domain.CardState, rating review_domain.RatingType, reviewDate time.Time) {
	// calculate the card's next interval
	// len(self.learning_steps) == 0: no learning steps defined so move card to Review state
	// card.step > len(self.learning_steps): handles the edge-case when a card was originally scheduled with a scheduler with more
	// learning steps than the current scheduler
	if len(s.learningSteps) == 0 || cardState.Step >= len(s.learningSteps) {
		cardState.
			SetState(card_state_domain.ReviewCardState).
			SetStep(0).
			SetEasy(s.startingEasy).
			SetCurrentIntervalInDays(s.graduationIntervalInDays).
			SetDue(reviewDate.Add(s.daysToTimeDuration(cardState.CurrentIntervalInDays)))

		return
	}

	switch rating {
	case review_domain.AgainRating:
		cardState.
			SetStep(0).
			SetDue(reviewDate.Add(s.learningSteps[cardState.Step]))
	case review_domain.HardRating:
		// card step stays the same
		if cardState.Step == 0 && len(s.learningSteps) == 1 {
			cardState.
				SetDue(reviewDate.Add(time.Duration(float64(s.learningSteps[cardState.Step].Nanoseconds()) * 1.5)))
		} else if cardState.Step == 0 && len(s.learningSteps) >= 2 {
			cardState.
				SetDue(reviewDate.Add(time.Duration(float64(s.learningSteps[cardState.Step]+s.learningSteps[cardState.Step+1]) * 0.5)))
		} else {
			cardState.
				SetDue(reviewDate.Add(s.learningSteps[cardState.Step]))
		}
	case review_domain.GoodRating:
		if cardState.Step+1 == len(s.learningSteps) { // the last step
			cardState.
				SetState(card_state_domain.ReviewCardState).
				SetStep(0).
				SetEasy(s.startingEasy).
				SetCurrentIntervalInDays(s.graduationIntervalInDays).
				SetDue(reviewDate.Add(s.daysToTimeDuration(s.graduationIntervalInDays)))
		} else {
			cardState.
				SetStep(cardState.Step + 1).
				SetDue(reviewDate.Add(s.learningSteps[cardState.Step]))
		}
	case review_domain.EasyRating:
		cardState.
			SetState(card_state_domain.ReviewCardState).
			SetStep(0).
			SetEasy(s.startingEasy).
			SetCurrentIntervalInDays(s.easyIntervalInDays).
			SetDue(reviewDate.Add(s.daysToTimeDuration(cardState.CurrentIntervalInDays)))
	}
}

func (s *Scheduler) stateReviewFunc(cardState *card_state_domain.CardState, rating review_domain.RatingType, reviewDate time.Time) {
	switch rating {
	case review_domain.AgainRating:
		cardState.SetEasy(max(1.3, cardState.Easy*0.8)) // Reduce easy by 20%

		currentIntervalInDays := max(s.minimalIntervalInDays, int64(math.Round(float64(cardState.CurrentIntervalInDays)*s.newInterval*s.intervalModifier)))
		cardState.SetCurrentIntervalInDays(s.fuzzy.get(currentIntervalInDays))

		// if there are no relearning steps (they were left blank)
		if len(s.relearningSteps) > 0 {
			cardState.
				SetState(card_state_domain.RelearningCardState).
				SetStep(0).
				SetDue(reviewDate.Add(s.relearningSteps[cardState.Step]))
		} else {
			cardState.SetDue(reviewDate.Add(s.daysToTimeDuration(cardState.CurrentIntervalInDays)))
		}
	case review_domain.HardRating:
		cardState.SetEasy(max(1.3, cardState.Easy*0.85)) // reduce ease by 15%

		currentIntervalIdDays := min(s.maximumIntervalInDays, int64(math.Round(float64(cardState.CurrentIntervalInDays)*s.hardInterval*s.intervalModifier)))
		cardState.
			SetCurrentIntervalInDays(s.fuzzy.get(currentIntervalIdDays)).
			SetDue(reviewDate.Add(s.daysToTimeDuration(currentIntervalIdDays)))
	case review_domain.GoodRating:
		// ease stays the same
		daysOverdue := (reviewDate.Sub(cardState.Due)).Nanoseconds() / (24 * time.Hour.Nanoseconds())

		currentIntervalInDays := int64(0)

		if daysOverdue >= 1 {
			currentIntervalInDays = min(s.maximumIntervalInDays, int64(math.Round((float64(cardState.CurrentIntervalInDays)+float64(daysOverdue)*0.5)*cardState.Easy*s.intervalModifier)))
		} else {
			currentIntervalInDays = min(s.maximumIntervalInDays, int64(math.Round(float64(cardState.CurrentIntervalInDays)*cardState.Easy*s.intervalModifier)))
		}

		cardState.
			SetCurrentIntervalInDays(s.fuzzy.get(currentIntervalInDays)).
			SetDue(reviewDate.Add(s.daysToTimeDuration(cardState.CurrentIntervalInDays)))
	case review_domain.EasyRating:
		daysOverdue := (reviewDate.Sub(cardState.Due)).Nanoseconds() / (24 * time.Hour.Nanoseconds())

		currentIntervalInDays := int64(0)

		if daysOverdue >= 1 {
			currentIntervalInDays = min(s.maximumIntervalInDays, int64(math.Round((float64(cardState.CurrentIntervalInDays)+float64(daysOverdue))*cardState.Easy*s.easyBonus*s.intervalModifier)))
		} else {
			currentIntervalInDays = min(s.maximumIntervalInDays, int64(math.Round(float64(cardState.CurrentIntervalInDays)*cardState.Easy*s.easyBonus*s.intervalModifier)))
		}

		cardState.
			SetCurrentIntervalInDays(s.fuzzy.get(currentIntervalInDays)).
			SetEasy(cardState.Easy * 1.15). // increase ease by 15%
			SetDue(reviewDate.Add(s.daysToTimeDuration(cardState.CurrentIntervalInDays)))
	}
}

func (s *Scheduler) stateRelearningFunc(cardState *card_state_domain.CardState, rating review_domain.RatingType, reviewDate time.Time) {
	// calculate the card's next interval
	// len(self.relearning_steps) == 0: no relearning steps defined so move card to Review state
	// card.step > len(self.relearning_steps): handles the edge-case when a card was originally scheduled with a scheduler with more
	// relearning steps than the current scheduler
	if len(s.relearningSteps) == 0 || cardState.Step >= len(s.relearningSteps) {
		cardState.SetState(card_state_domain.ReviewCardState).SetStep(0)
		// don't update ease
		cardState.
			SetCurrentIntervalInDays(min(s.maximumIntervalInDays, int64(math.Round(float64(cardState.CurrentIntervalInDays)*cardState.Easy*s.intervalModifier)))).
			SetDue(reviewDate.Add(s.daysToTimeDuration(cardState.CurrentIntervalInDays)))

		return
	}

	switch rating {
	case review_domain.AgainRating:
		cardState.SetStep(0).SetDue(reviewDate.Add(s.relearningSteps[cardState.Step]))
	case review_domain.HardRating:
		// cardState step stays the same
		if cardState.Step == 0 && len(s.relearningSteps) == 1 {
			cardState.SetDue(reviewDate.Add(time.Duration(float64(s.relearningSteps[cardState.Step].Nanoseconds()) * 1.5)))
		} else if cardState.Step == 0 && len(s.relearningSteps) >= 2 {
			cardState.SetDue(reviewDate.Add(time.Duration(float64(s.relearningSteps[cardState.Step]+s.relearningSteps[cardState.Step+1]) * 0.5)))
		} else {
			cardState.SetDue(reviewDate.Add(s.relearningSteps[cardState.Step]))
		}
	case review_domain.GoodRating:
		if cardState.Step+1 == len(s.relearningSteps) { // the last step
			cardState.SetState(card_state_domain.ReviewCardState).SetStep(0)
			// don't update ease
			cardState.
				SetCurrentIntervalInDays(min(s.maximumIntervalInDays, int64(math.Round(float64(cardState.CurrentIntervalInDays)*cardState.Easy*s.intervalModifier)))).
				SetDue(reviewDate.Add(s.daysToTimeDuration(cardState.CurrentIntervalInDays)))
		} else {
			cardState.
				SetStep(cardState.Step + 1).
				SetDue(reviewDate.Add(s.relearningSteps[cardState.Step]))
		}
	case review_domain.EasyRating:
		cardState.SetState(card_state_domain.ReviewCardState).SetStep(0)
		// don't update ease
		cardState.
			SetCurrentIntervalInDays(min(s.maximumIntervalInDays, int64(math.Round(float64(cardState.CurrentIntervalInDays)*cardState.Easy*s.easyBonus*s.intervalModifier)))).
			SetDue(reviewDate.Add(s.daysToTimeDuration(cardState.CurrentIntervalInDays)))
	}
}

func (s *Scheduler) daysToTimeDuration(days int64) time.Duration {
	return time.Duration(days * 24 * time.Hour.Nanoseconds())
}
