package myremo

import (
	"context"
	"errors"
	"log"

	"github.com/tenntenn/natureremo"
)

// Remo is struct for communicating to Natureremo API
type Remo struct {
	Client    *natureremo.Client
	Signal    *natureremo.Signal
	Appliance *natureremo.Appliance
}

// NewRemo is contstructor for Nature Remo API
func NewRemo(token string) *Remo {
	remo := new(Remo)
	remo.Client = natureremo.NewClient(token)
	return remo
}

// GetAppliance is function to getting all of appliances
func (r *Remo) GetAppliance(ctx context.Context, name string) error {
	appliances, err := r.Client.ApplianceService.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, a := range appliances {
		if a.Nickname == name {
			r.Appliance = a
			return nil
		}
	}

	return errors.New("appliance not found")
}

// GetSignal is function to getting Signal
func (r *Remo) GetSignal(ss []*natureremo.Signal, name string) error {
	for _, s := range ss {
		if s.Name == name {
			r.Signal = s
			return nil
		}
	}
	return errors.New("Signal Not Found.")
}

// SendSignal is function to send signal to remo API
func (r *Remo) SendSignal(ctx context.Context) error {
	if err := r.Client.SignalService.Send(ctx, r.Signal); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// SendSignalByAplSig is function to send signal by appliance and signal
func (r *Remo) SendSignalByAplSig(apl, sig string, ctx context.Context) error {
	if err := r.GetAppliance(ctx, apl); err != nil {
		log.Fatal(err)
		return err
	}
	if err := r.GetSignal(r.Appliance.Signals, sig); err != nil {
		log.Fatal(err)
		return err
	}
	if err := r.SendSignal(ctx); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
