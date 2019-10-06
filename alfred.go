package Alfred

import (
	"encoding/json"
	"bytes"
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
	"howett.net/plist"
)

const (
	plistInfoFileName string = "info.plist"
	defaultCacheDir   string = "Library/Caches/com.runningwithcrayons.Alfred/Workflow Data"
	defaultDataDir    string = "Library/Application Support/Alfred/Workflow Data"
)

type Alfred struct {
	id        string
	homeDir   string
	bundleDir string
	bundleID  string
	dataDir   string
	cacheDir  string
	result    *Res
}


func (al *Alfred) init(id string) (*Alfred, error) {
	al.id = id

	err := al.initHomeDir()
	if err != nil {
		return nil, errors.Wrap(err, sign("init homedir fail"))
	}

	err = al.initBundle()
	if err != nil {
		return nil, errors.Wrap(err, sign("init bundledir fail"))
	}

	err = al.initBundleID()
	if err != nil {
		return nil, errors.Wrap(err, sign("init bundle id fail"))
	}

	al.cacheDir = path.Join(al.homeDir, defaultCacheDir, al.bundleID)
	al.dataDir = path.Join(al.homeDir, defaultDataDir, al.bundleID)

	err = al.initCacheDir()
	if err != nil {
		return nil, errors.Wrap(err, sign("init cache dir fail"))
	}

	err = al.initDataDir()
	if err != nil {
		return nil, errors.Wrap(err, sign("init data dir fail"))
	}

	al.result = new(Res)

	return al, nil
}

func (al *Alfred) initHomeDir() error {
	home := os.Getenv("HOME")

	if home == "" {
		return errors.New(sign("find home dir fail"))
	}
	al.homeDir = home
	return nil
}

func (al *Alfred) initBundle() error {
	pwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, sign("get pwd fail"))
	}
	al.bundleDir = pwd

	return nil
}

func (al *Alfred) initBundleID() error {
	// Find info.plist
	plistFile := path.Join(al.bundleDir, plistInfoFileName)
	_, err := os.Stat(plistFile)
	if err != nil {
		return errors.Wrap(err, sign("find info.plist fail", plistFile))
	}

	// Read info.plist
	buf, err := ioutil.ReadFile(plistFile)
	if err != nil {
		return errors.Wrap(err, sign("read info.plist fail", plistFile))
	}

	// Parse info.plist
	var attrs map[string]interface{}
	decoder := plist.NewDecoder(bytes.NewReader(buf))
	err = decoder.Decode(&attrs)
	if err != nil {
		return errors.Wrap(err, sign("parse info.plist fail", plistFile))
	}

	v, ok := attrs["bundleid"]
	if !ok {
		return errors.Wrap(err, sign("find bundle id fail", plistFile))
	}

	al.bundleID = v.(string)
	return nil
}

func (al *Alfred) initCacheDir() error {
	if _, err := os.Stat(al.cacheDir); err != nil {
		if err = os.MkdirAll(al.cacheDir, 0755); err != nil {
			return errors.Wrap(err, sign("create cache dir fail", al.cacheDir))
		}
	}

	return nil
}

func (al *Alfred) initDataDir() error {
	if _, err := os.Stat(al.dataDir); err != nil {
		if err = os.MkdirAll(al.dataDir, 0755); err != nil {
			return errors.Wrap(err, sign("create data dir fail", al.dataDir))
		}
	}

	return nil
}

func (al *Alfred) SetCacheDir(dir string) (*Alfred, error) {
	al.cacheDir = dir
	if err := al.initCacheDir(); err != nil {
		return nil, err
	}

	return al, nil
}

func (al *Alfred) SetDataDir(dir string) (*Alfred, error) {
	al.dataDir = dir
	if err := al.initDataDir(); err != nil {
		return nil, err
	}

	return al, nil
}

func (al *Alfred) GetID() string {
	return al.id
}

func (al *Alfred) GetHomeDir() string {
	return al.homeDir
}

func (al *Alfred) GetBundleDir() string {
	return al.bundleDir
}

func (al *Alfred) GetBundleID() string {
	return al.bundleID
}

func (al *Alfred) GetDataDir() string {
	return al.dataDir
}

func (al *Alfred) GetCacheDir() string {
	return al.cacheDir
}

func (al *Alfred) GetResult() *Res{
	return al.result
}

func (al *Alfred) ResultAppend(it Item) *Alfred{
	al.result.Append(it)
	return al
}

func (al *Alfred) ResultSet(its []Item) *Alfred{
	al.result.Set(its)
	return al
}

func (al *Alfred) ResultToJson() ([]byte, error){
	return json.Marshal(al.result)
}

func (al *Alfred) ResultToIndentJson() ([]byte, error){
	return json.MarshalIndent(al.result, "", "  ")
}

func (al *Alfred) Output() (int, error){
	if len(al.result.Items) == 0 {
		al.ResultAppend(*NewNoResultItem())
	}

	json, err := al.ResultToJson()
	if err != nil {
		return os.Stdout.WriteString("{\"items\":[{\"uid\":\"\",\"type\":\"default\",\"title\":\"We had a error\",\"subtitle\":\"~/Desktop\",\"arg\":\"\",\"autocomplete\":\"\",\"valid\":false,\"icon\":{\"type\":\"fileicon\",\"path\":\"~/Desktop\"}}]}")
	}
	
	return os.Stdout.Write(json)
}


func NewAlfred(id string) (*Alfred, error) {
	return new(Alfred).init(id)
}
