/**
 * Author:        StEvUgnIn <steve.huguenin-elie@protonmail.ch>
 * SPDX License:  AGPL-3.0
 * Description:   This controller is based on [go-orbit-db](berty.tech/go-orbit-db) to simplify working with OrbitDB.
 */
package db

import (
	"context"
	"encoding/base32"
	"log"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/address"
	"berty.tech/go-orbit-db/iface"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	coreapi "github.com/ipfs/interface-go-ipfs-core"
)

const (
	KeyValueStore = "keyvalue"
	DocumentStore = "docstore"
	EventLogStore = "eventlog"
)

func DefaultCreateDBOptions() orbitdb.CreateDBOptions {
	Replicate, LocalOnly := boolPtr(true), boolPtr(false)
	if version == "Development" {
		Replicate, LocalOnly = LocalOnly, Replicate
	}
	return orbitdb.CreateDBOptions{
		Replicate: Replicate,
		LocalOnly: LocalOnly,
	}
}

func NewLocalController(ctx context.Context, name string, storeType string, dbOpts iface.CreateDBOptions, options *orbitdb.NewOrbitDBOptions) (*OrbitController, error) {
	i, err := httpapi.NewLocalApi()
	if err != nil {
		return nil, err
	}
	return NewController(ctx, i, name, storeType, dbOpts, options)
}

func NewController(ctx context.Context, i coreapi.CoreAPI, name string, storeType string, dbOpts iface.CreateDBOptions, options *orbitdb.NewOrbitDBOptions) (*OrbitController, error) {
	if dbOpts.Replicate == nil {
		dbOpts.Replicate = DefaultCreateDBOptions().Replicate
	}
	if dbOpts.LocalOnly == nil {
		dbOpts.LocalOnly = DefaultCreateDBOptions().LocalOnly
	}
	db, err := orbitdb.NewOrbitDB(ctx, i, options)
	if err != nil {
		return nil, err
	}
	err = address.IsValid(name)
	var addr address.Address

	if err == nil {
		addr, err = db.DetermineAddress(ctx, name, storeType, &iface.DetermineAddressOptions{
			OnlyHash:  boolPtr(base32.StdEncoding.EncodeToString([]byte(name)) == name),
			Replicate: dbOpts.Replicate,
		})
	}
	var store iface.Store
	if err == nil {
		store, err = db.Open(ctx, addr.String(), &dbOpts)
	} else {
		store, err = db.Create(ctx, name, storeType, &dbOpts)
	}
	if err != nil {
		return nil, err
	}
	return &OrbitController{
		ctx:    ctx,
		db:     db,
		store:  store,
		dbOpts: &dbOpts,
	}, nil
}

type OrbitController struct {
	ctx    context.Context
	db     iface.OrbitDB
	dbOpts *iface.CreateDBOptions
	store  iface.Store
}

type Controller interface {
	Ctx() context.Context
	OrbitDB() iface.OrbitDB
	CreateDBOptions() *iface.CreateDBOptions
	Store() iface.Store
	KeyValueStore() (iface.KeyValueStore, error)
	DocumentStore() (iface.DocumentStore, error)
	EventLogStore() (iface.EventLogStore, error)
}

func (controller *OrbitController) Ctx() context.Context {
	return controller.ctx
}

func (controller *OrbitController) OrbitDB() iface.OrbitDB {
	return controller.db
}

func (controller *OrbitController) CreateDBOptions() *iface.CreateDBOptions {
	return controller.dbOpts
}

func (controller *OrbitController) Store() iface.Store {
	return controller.store
}

func (controller *OrbitController) KeyValueStore() (iface.KeyValueStore, error) {
	if controller.store.Type() != KeyValueStore {
		return controller.db.KeyValue(controller.ctx, controller.store.Address().String(), controller.dbOpts)
	}
	return controller.store.(iface.KeyValueStore), nil
}

func (controller *OrbitController) DocumentStore() (iface.DocumentStore, error) {
	if controller.store.Type() != DocumentStore {
		return controller.db.Docs(controller.ctx, controller.store.Address().String(), controller.dbOpts)
	}
	return controller.store.(iface.DocumentStore), nil
}

func (controller *OrbitController) EventLogStore() (iface.EventLogStore, error) {
	if controller.store.Type() != EventLogStore {
		return controller.db.Log(controller.ctx, controller.store.Address().String(), controller.dbOpts)
	}
	return controller.store.(iface.EventLogStore), nil
}

func Must(controller *OrbitController, err error) *OrbitController {
	if err != nil {
		log.Fatalln(err)
	}
	return controller
}
