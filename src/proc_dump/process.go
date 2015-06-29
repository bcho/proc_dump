package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	lp "github.com/c9s/goprocinfo/linux"
)

type ProcessMap map[string]Process

func parseProcs(raw []string) (ProcessMap, error) {
	var (
		procs = make(ProcessMap)
		proc  *Process
		err   error
	)
	for _, pid := range raw {
		_, err = strconv.ParseUint(pid, 10, 64)
		if err != nil {
			return nil, err
		}

		proc, err = ReadProcess(pid, ProcPath)
		if err != nil {
			return nil, err
		}
		procs[pid] = *proc
	}
	return procs, nil
}

func (p ProcessMap) Dump() ([]byte, error) {
	return json.Marshal(p)
}

func (p ProcessMap) Refresh() error {
	for _, proc := range p {
		if err := proc.Refresh(); err != nil {
			return err
		}
	}
	return nil
}

type Process struct {
	pid  string `json:"-"`
	path string `json:"-"`

	Stat lp.ProcessStat `json:"stat"`
}

func (p *Process) Refresh() (err error) {
	procPath := filepath.Join(p.path, p.pid)
	if _, err = os.Stat(procPath); err != nil {
		return err
	}

	stat, err := lp.ReadProcessStat(filepath.Join(procPath, "stat"))
	if err != nil {
		return err
	}
	p.Stat = *stat

	return nil
}

func ReadProcess(pid, path string) (*Process, error) {
	proc := Process{pid: pid, path: path}
	if err := proc.Refresh(); err != nil {
		return nil, err
	}

	return &proc, nil
}
