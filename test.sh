#!/bin/bash
while ! ~/robinAI/robin
do
	echo "restarting"
	exec ./robin
done

