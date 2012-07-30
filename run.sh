#!/bin/sh

ruby ./times_to_graph.rb

# this verison sets a switching cost of 10 each time a line changes
./buildmatrix -s 10

# this version has switching cost set to 0 (no penalty for switching lines)
#./buildmatrix

./affectation
./buildmatrix -p
./dump_route.rb