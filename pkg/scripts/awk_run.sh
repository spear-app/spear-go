#!/bin/bash
awk '{print $4,$8}' $1 > $2
