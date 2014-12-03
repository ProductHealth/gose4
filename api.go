package gose4

import "github.com/golang/glog"

// Logging wrappers
var Infof = glog.Infof
var Warningf = glog.Warningf
var Debugf = glog.V(2).Infof
