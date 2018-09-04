// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"

	"github.com/medibloc/go-medibloc/medlet"
	log "github.com/medibloc/go-medibloc/util/logging"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	version string
	commit  string
	branch  string
)

func main() {
	app := cli.NewApp()
	app.Action = medi
	app.Name = "medi"
	app.Usage = "medibloc command line interface"
	app.Version = versionStr()

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()

	app.Run(os.Args)
}
func versionStr() string {
	if version == "" {
		return ""
	}
	return fmt.Sprintf("%s, branch %s, commit %s", version, branch, commit)
}

func medi(ctx *cli.Context) error {
	configFile := ctx.Args().Get(0)
	conf, err := medlet.LoadConfig(configFile)
	if err != nil {
		log.Console().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to load config.")
		return err
	}
	m, err := medlet.New(conf)
	if err != nil {
		log.Console().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to create medlet.")
		return err
	}

	log.Init(m.Config().App.LogFile, m.Config().App.LogLevel, m.Config().App.LogAge)

	return runMedi(ctx, m)
}

func runMedi(ctx *cli.Context, m *medlet.Medlet) error {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, os.Kill)

	// Run Node
	log.Console().Info("Start medibloc...")

	err := m.Setup()
	if err != nil {
		log.Console().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to setup medlet.")
		return err
	}

	err = m.Start()
	if err != nil {
		log.Console().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to start medlet.")
		return err
	}

	<-sigch
	m.Stop()
	log.Console().Info("Stop medibloc...")
	return nil
}
