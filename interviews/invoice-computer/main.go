package main

import (
	"fmt"
	"strconv"
	"time"
)

type Subscription struct {
	Id                    int
	CustomerId            int
	MonthlyPriceInDollars int
}

type User struct {
	Id            int
	Name          string
	ActivatedOn   time.Time
	DeactivatedOn time.Time
	CustomerId    int
}

// Computes the monthly charge for a given subscription.
//
// Returns the total monthly bill for the customer in dollars and cents, rounded
// to 2 decimal places.
// If there are no active users or the subscription is null, returns 0.
//
// month: Always present
//
//	Has the following structure:
//	"2022-04"  // April 2022 in YYYY-MM format
//
// activeSubscription: May be null
//
//	If present, has the following structure (see Subscription struct):
//	{
//	  Id: 763,
//	  CustomerId: 328,
//	  MonthlyPriceInDollars: 4  // price per active user per month
//	}
//
// users: May be empty, but not null
//
//	Has the following structure (see User struct):
//	{
//	  {
//	    Id: 1,
//	    Name: "Employee #1",
//	    CustomerId: 1,
//
//	    // when this user started
//	    Activated_On: time.Date(2021, 11, 4, 0, 0, 0, 0, time.UTC),
//
//	    // last day to bill for user
//	    // should bill up to and including this date
//	    // since user had some access on this date
//	    DeactivatedOn: time.Date(2022, 4, 10, 0, 0, 0, 0, time.UTC)
//	  },
//	  {
//	    Id: 2,
//	    Name: "Employee #2",
//	    CustomerId: 1,
//
//	    // when this user started
//	    ActivatedOn: time.Date(2021, 12, 4, 0, 0, 0, 0, time.UTC),
//
//	    // hasn't been deactivated yet
//	    DeactivatedOn: nil
//	  },
//	}

// computeDailyRate computes the daily rate of the given subscription. Month
// is given in time format, out of which the month is the relevant aspect.
func computeDailyRate(s *Subscription, month time.Time) float64 {
	// Adding 1 since it's an inclusive interval.
	totalDays := LastDayOfMonth(month).Day() - FirstDayOfMonth(month).Day() + 1
	return float64(s.MonthlyPriceInDollars) / float64(totalDays)
}

// isUserActive returns whether the user is active or not in that day.
func isUserActive(u *User, day time.Time) bool {

	// For the purpose of this example I will treat this as an anomaly.
	// Obviously it's just a shortcut here...
	if !u.ActivatedOn.IsZero() && !u.DeactivatedOn.IsZero() && u.ActivatedOn.After(u.DeactivatedOn) {
		panic("strange case encountered")
	}

	// Whether or not the account was activa
	isActivatedBefore := u.ActivatedOn.IsZero() || u.ActivatedOn.Before(NextDay(day))
	isDeactivatedBefore := !u.DeactivatedOn.IsZero() && u.DeactivatedOn.Before(day)

	// The account has to be activated before (or during) that day, and also not
	// be deactivated before the start of the day.
	return isActivatedBefore && !isDeactivatedBefore
}

func BillFor(yearMonth string, activeSubscription *Subscription, users *[]User) float64 {

	// edge cases
	if activeSubscription == nil {
		return 0
	}
	if users == nil {
		return 0
	} else {
		if len(*users) == 0 {
			return 0
		}
	}

	currentBillingMonth, err := time.Parse("2006-01", yearMonth)
	if err != nil {
		// It is bad practice to panic in production. This function does not
		// include an error in the signature, but normally you would return the
		// error here.
		panic(err)
	}

	dailyRate := computeDailyRate(activeSubscription, currentBillingMonth)
	var ongoingTotal float64 = 0

	// currentDay holds the start of the given day.
	currentDay := FirstDayOfMonth(currentBillingMonth)

	// We care for the days that are inside that month (thus, before the
	// first day of the next month). Note that LastDayOfMonth is the last
	// possible time value of that month.
	for currentDay.Before(LastDayOfMonth(currentBillingMonth)) {
		fmt.Println(currentDay.Day())
		noActiveUsers := 0
		for _, u := range *users {
			if isUserActive(&u, currentDay) {
				noActiveUsers++
			}
		}
		ongoingTotal += float64(noActiveUsers) * float64(dailyRate)
		currentDay = NextDay(currentDay)
	}

	// Total amount, rounded to 2 decimal places.

	ongoingTotal, err = strconv.ParseFloat(fmt.Sprintf("%.2f", ongoingTotal), 64)
	if err != nil {
		// Again shortcut...
		panic(err)
	}
	return ongoingTotal
}

// func BillFor(yearMonth string, activeSubscription *Subscription, users *[]User) float64 {
// 	// If there are no active users or the subscription is null, returns 0.
// 	// This covers basic edge cases (note that float is returned automatically).
// 	if activeSubscription == nil {
// 		return 0
// 	}
// 	if users == nil {
// 		return 0
// 	} else {
// 		if len(*users) == 0 {
// 			return 0
// 		}
// 	}

// 	currentBillingMonth, err := time.Parse("2006-01", yearMonth)
// 	if err != nil {
// 		// It is bad practice to panic in production. This function does not
// 		// include an error in the signature, but normally you would return the
// 		// error here.
// 		panic(err)
// 	}

// 	// FirstDay and LastDay are required to compute whether or not the user's
// 	// subscription period intersects with the ongoing month.
// 	firstDay := FirstDayOfMonth(currentBillingMonth)
// 	lastDay := LastDayOfMonth(currentBillingMonth)

// 	// Note that firstDay and lastDay are included in the month.
// 	monthDays := lastDay.Day() - firstDay.Day() + 1
// 	var totalAmount float64 = 0

// 	// In order to optimize the floating point arithmetic, we will keep track
// 	// of the total days that are billed, rather than computing the total
// 	// amount each user has to pay (which is a floating point) and add those
// 	// at the end.
// 	var totalDaysToBillFor int

// 	// Find the total number of days users were active, cummulated.
// 	for _, user := range *users {
// 		// These are intermediate variables to store when the users are active
// 		// during the ongoing month. If the subscription started before the
// 		// month, it will store the first day of the month, and last day of the
// 		// month if the subscription ends after the month completes.
// 		var activatedDay, deactivatedDay int
// 		// FirstDay has all fields expect day, month and year set to 0, so there
// 		// is no need to decrease.
// 		if user.ActivatedOn.Equal(time.Time{}) || user.ActivatedOn.Before(firstDay) {
// 			activatedDay = firstDay.Day()
// 		} else {
// 			activatedDay = user.ActivatedOn.Day()
// 		}

// 		// Use the next day in order to account for the users that end their
// 		// subscription in the last day. Like above, all fields of next day
// 		// except year, month and day are 0 so this is valid.
// 		if user.DeactivatedOn.Equal(time.Time{}) || user.DeactivatedOn.After(NextDay(lastDay)) {
// 			deactivatedDay = lastDay.Day()
// 		} else {
// 			deactivatedDay = user.DeactivatedOn.Day()
// 		}

// 		// Compute the number of days the user was active (inclusive).
// 		totalUserDays := deactivatedDay - activatedDay + 1
// 		totalDaysToBillFor += totalUserDays
// 	}

// 	// Total amount, computed using the total number of billable days.
// 	totalAmount += float64(totalDaysToBillFor) * float64(activeSubscription.MonthlyPriceInDollars) / float64(monthDays)

// 	//totalAmount, err = strconv.ParseFloat(fmt.Sprintf("%2f", totalAmount), 64)
// 	return totalAmount
// }

/*******************
* Helper functions *
*******************/

/*
Takes a time.Time object and returns a time.Time object
which is the first day of that month.

FirstDayOfMonth(time.Date(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb 7
=> time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC))               // Feb 1
*/
func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

/*
Takes a time.Time object and returns a time.Time object
which is the end of the last day of that month.

LastDayOfMonth(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb  7
=> time.Time(2019, 2, 28, 23, 59, 59, 0, time.UTC)           // Feb 28
*/
func LastDayOfMonth(t time.Time) time.Time {
	return FirstDayOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

/*
Takes a time.Time object and returns a time.Time object
which is the next day.

NextDay(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))   // Feb 7
=> time.Time(2019, 2, 8, 0, 0, 0, 0, time.UTC)         // Feb 8

NextDay(time.Time(2019, 2, 28, 0, 0, 0, 0, time.UTC))  // Feb 28
=> time.Time(2019, 3, 1, 0, 0, 0, 0, time.UTC)         // Mar  1
*/
func NextDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func main() {
	constantUsers := []User{
		{
			Id:            1,
			Name:          "Employee #1",
			ActivatedOn:   time.Date(2018, 11, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            2,
			Name:          "Employee #2",
			ActivatedOn:   time.Date(2018, 12, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
	}

	userSignedUp := []User{
		{
			Id:            1,
			Name:          "Employee #1",
			ActivatedOn:   time.Date(2018, 11, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            2,
			Name:          "Employee #2",
			ActivatedOn:   time.Date(2018, 12, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            3,
			Name:          "Employee #3",
			ActivatedOn:   time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
	}

	newPlan := Subscription{
		Id:                    1,
		CustomerId:            1,
		MonthlyPriceInDollars: 4,
	}

	var noUsers []User

	fmt.Println("test 1:", BillFor("2019-01", &newPlan, &noUsers))
	fmt.Println("test 2:", BillFor("2019-01", &newPlan, &constantUsers))
	fmt.Println("test 3:", BillFor("2019-01", &newPlan, &userSignedUp))
}
