package controllers

import "fmt"

type State interface {
	RequestSub(sub bool) error
	InsertMoney(money int) error
	SuccessSub() error
}

type Subs struct {
	CheckSub   State
	IsSub      State
	IsNotSub   State
	CheckMoney State
	HasMoney   State

	CurrentState State

	itemPrice int
}

func NewSubs(itemPrice int) *Subs {
	v := &Subs{
		itemPrice: itemPrice,
	}
	CheckSubState := &CheckSubState{
		Subs: v,
	}
	isSubState := &isSubState{
		Subs: v,
	}
	isNotSubState := &isNotSubState{
		Subs: v,
	}
	checkMoneyState := &checkMoneyState{
		Subs: v,
	}
	hasMoneyState := &hasMoneyState{
		Subs: v,
	}

	v.SetState(CheckSubState)
	v.CheckSub = CheckSubState
	v.IsSub = isSubState
	v.IsNotSub = isNotSubState
	v.CheckMoney = checkMoneyState
	v.HasMoney = hasMoneyState
	return v
}

func (v *Subs) RequestSub(sub bool) error {
	return v.CurrentState.RequestSub(sub)
}

func (v *Subs) InsertMoney(money int) error {
	return v.CurrentState.InsertMoney(money)
}

func (v *Subs) SuccessSubs() error {
	return v.CurrentState.SuccessSub()
}

func (v *Subs) SetState(s State) {
	v.CurrentState = s
}

type CheckSubState struct {
	Subs *Subs
}

func (i *CheckSubState) RequestSub(sub bool) error {
	if sub {
		i.Subs.SetState(i.Subs.IsSub)
	} else {
		i.Subs.SetState(i.Subs.IsNotSub)
	}
	return nil
}

func (i *CheckSubState) InsertMoney(money int) error {
	return fmt.Errorf("Checking account")
}
func (i *CheckSubState) SuccessSub() error {
	return fmt.Errorf("Checking account")
}

type isSubState struct {
	Subs *Subs
}

func (i *isSubState) RequestSub(sub bool) error {
	return fmt.Errorf("You are already sub")
}

func (i *isSubState) InsertMoney(money int) error {
	return fmt.Errorf("You are already sub")
}
func (i *isSubState) SuccessSub() error {
	return fmt.Errorf("You are already sub")
}

type isNotSubState struct {
	Subs *Subs
}

func (i *isNotSubState) RequestSub(sub bool) error {
	fmt.Println("You are not sub")
	i.Subs.SetState(i.Subs.CheckMoney)
	return nil
}

func (i *isNotSubState) InsertMoney(money int) error {
	i.Subs.SetState(i.Subs.CheckMoney)
	return nil
}
func (i *isNotSubState) SuccessSub() error {
	return fmt.Errorf("Please wait")
}

type checkMoneyState struct {
	Subs *Subs
}

func (i *checkMoneyState) RequestSub(sub bool) error {
	return fmt.Errorf("Please wait")
}

func (i *checkMoneyState) InsertMoney(money int) error {
	if money < i.Subs.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.Subs.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.Subs.SetState(i.Subs.HasMoney)
	return nil
}
func (i *checkMoneyState) SuccessSub() error {
	fmt.Println("To fuck")
	return nil
}

type hasMoneyState struct {
	Subs *Subs
}

func (i *hasMoneyState) RequestSub(sub bool) error {
	return fmt.Errorf("Please wait")
}

func (i *hasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Please wait")
}
func (i *hasMoneyState) SuccessSub() error {
	fmt.Println("Successfully subscribed")
	i.Subs.SetState(i.Subs.IsSub)
	return nil
}
