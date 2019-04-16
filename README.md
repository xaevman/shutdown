# Shutdown

## Build Status
[![Build Status](https://travis-ci.org/xaevman/shutdown.svg?branch=master)](https://travis-ci.org/xaevman/shutdown)

## Overview
Shutdown is a Go package which aims to help provide clean and deterministic behavior during application shutdown. Shutdown sync objects allow you to communicate the intent to shutdown, 
execute shutdown code, signal when shutdown is complete, and also provide for timeouts and error handling should shutdown procedures hang, deadlock, or do something unexpected.

### License
This Go package is release under a BSD-style license, the contents of which are in the repo's LICENSE file.

## API Documentation
http://godoc.org/github.com/xaevman/shutdown
