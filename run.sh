#!/bin/sh

ruby ./times_to_graph.rb
./buildmatrix
./affectation
./buildmatrix -p
./dump_route.rb