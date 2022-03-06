#!/bin/bash
#!/usr/bin/env bash

echo "Start REST API Server:"

command objy StartRESTServer -fdAlias kb -configFile restApiConfig.xml
