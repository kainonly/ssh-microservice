package utils

import (
	"github.com/robfig/cron/v3"
)

type CronMap struct {
	hashMap map[string]*cron.Cron
}

func NewCronMap() *CronMap {
	c := new(CronMap)
	c.hashMap = make(map[string]*cron.Cron)
	return c
}

func (c *CronMap) Put(identity string, value *cron.Cron) {
	c.hashMap[identity] = value
}

func (c *CronMap) Empty(identity string) bool {
	return c.hashMap[identity] == nil
}

func (c *CronMap) Get(identity string) *cron.Cron {
	return c.hashMap[identity]
}

func (c *CronMap) Remove(identity string) {
	delete(c.hashMap, identity)
}
