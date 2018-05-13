#!/bin/bash

if [ "$#" -eq 0 ]; then
  echo "No name entered."
else
  names=( "$@" )
  javac SoundexConverter.java
  java SoundexConverter "$@"
fi