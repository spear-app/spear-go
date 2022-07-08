#!/bin/bash
awk '{print $4}' $1 > $2
awk '{print $8}' $1 > $3
