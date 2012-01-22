#!/usr/bin/python
#-*- coding:utf-8 -*-

class Node:
	
	def __init__(self, name, id):
		self.name = name
		self.id=id
		
	def setCoord(self, x, y):
		self.x=x
		self.y=y


nodes = []
edges={}
with open("../ressources/metro.graph") as f:
	[nbnodes, nbedges] = [int(s) for s in f.readline().split(" ")]
	f.readline()
	for i in range(nbnodes):
		nodes.append( Node(" ".join(f.readline().split(" ")[1:]), i))
	f.readline()
	for i in range(nbnodes):
		[x,y] = [int(s) for s in f.readline().split(" ")][1:]
		nodes[i].setCoord(x, y)
	f.readline()
	for i in range(nbedges):
		[start, end, cost] = [int(s.strip()) for s in f.readline().split(" ")]
		if start in edges:
			edges[start].append((end, cost))
		else:
			edges[start]= [(end, cost)]
	k=0
	for i in edges:
		if(len(edges[i])==1):
			k+=1
			print nodes[i].name
	print k
