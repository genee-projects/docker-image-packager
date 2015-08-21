#!/bin/bash

for i in *.tar; do
        echo "$i"
		docker load -i "$i"
done

echo "done"
