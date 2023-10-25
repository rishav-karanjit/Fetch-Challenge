package main

import "sync"

var pointsDB = map[string]int64{}
var pointsMutex = &sync.Mutex{}

var receiptDB = map[string]Receipt{}
var receiptMutex = &sync.Mutex{}