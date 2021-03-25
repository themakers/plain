package storage

import (
    "bytes"
    "fmt"
    "github.com/themakers/plain/internal/state/state_v1"
    "io/ioutil"
    "log"
    "os"
)

const defaultStorePath = "~/.plain.yaml"

type Storage struct {
	oldState []byte
}

func New() *Storage { return &Storage{} }

func (s *Storage) loadRaw() []byte {
	path := shellExpand(defaultStorePath)

	if file, err := os.Open(path); err != nil {
		if os.IsNotExist(err) {
			return mustMarshal(newInitialState())
		} else {
			panic(err)
		}
	} else {
		defer file.Close()
		if data, err := ioutil.ReadAll(file); err != nil {
			panic(err)
		} else {
			return data
		}
	}
}

func (s *Storage) Load() *state_v1.State {
	if len(s.oldState) != 0 {
		panic("state is already loaded")
	}

	data := s.loadRaw()

	s.oldState = data

	state := mustUnmarshal(data)

	if state.Version != "1" {
		panic(fmt.Sprintf("unsupported state version %s, please update the app", state.Version))
	}

	return state
}

func (s *Storage) Save(state *state_v1.State) {
	diskData := s.loadRaw()
	if bytes.Compare(diskData, s.oldState) != 0 {
		panic("state file was modified on disk")
	}

	var (
		path    = shellExpand(defaultStorePath)
		pathBcp = fmt.Sprintf("%s-bak", path)
		pathTmp = fmt.Sprintf("%s-tmp", path)
		data    = mustMarshal(state)
	)

	if bytes.Compare(data, s.oldState) == 0 {
		log.Println("WARNING", "new and cached states are identical")
		// return
		// panic("new and cached states are identical")
	} else if bytes.Compare(data, diskData) == 0 {
		panic("new and disk states are identical")
	}

	if err := ioutil.WriteFile(pathTmp, data, 0600); err != nil {
		panic(err)
	} else if err := os.Remove(pathBcp); err != nil && !os.IsNotExist(err) {
		panic(err)
	} else if err := os.Rename(path, pathBcp); err != nil && !os.IsNotExist(err) {
		panic(err)
	} else if err := os.Rename(pathTmp, path); err != nil {
		panic(err)
	}

	s.oldState = data

	log.Println("SAVED")
}
