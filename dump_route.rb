#!/usr/bin/env ruby

# Expects a list of connected nodes in routes (connections should be in edges.csv).
# For any discontinuous jumps, the jump is printed and the calculation is picked up at next location.
#route = %w(170 169 168 167 161 162 163 164 165 166 165 164 163 162 161 160 116 159 158 157 156 155 154 10 153 10 11 10 9 8 7 191 190 189 55 108 82 147 146 147 82 148 149 150 151 63 64 65 42 66 67 68 69 70 71 72 71 70 69 68 67 66 42 41 40 286 287 288 289 290 291 292 291 290 289 288 287 286 293 294 295 296 297 298 299 300 299 298 297 296 295 294 293 286 40 39 38 37 36 35 34 107 175 174 173 176 172 171 173 176 252 253 252 176 251 74 250 74 73 75 51 49 50 49 51 52 53 54 55 56 57 58 59 60 61 62 152 12 13 14 267 266 265 243 264 263 93 262 93 128 95 96 97 98 99 98 97 96 95 94 93 92 91 90 89 244 243 242 241 283 284 184 15 221 222 221 220 16 17 18 19 20 21 22 23 24 23 22 21 20 19 18 17 16 220 221 64 63 62 186 187 188 84 55 254 31 253 252 176 172 171 173 174 175 107 34 33 32 31 254 55 56 57 255 57 58 84 188 187 186 62 185 14 184 183 182 132 181 180 179 177 178 177 179 180 181 132 239 238 237 236 213 212 211 210 209 207 208 207 209 210 211 212 213 235 233 234 233 212 248 249 237 236 213 214 215 216 217 136 218 219 218 136 137 138 137 136 135 134 133 132 131 130 129 261 260 259 258 256 257 256 258 259 260 261 129 130 131 132 239 238 239 132 240 241 242 243 244 89 245 246 247 246 245 89 88 87 10 86 85 58 84 83 82 81 80 79 78 76 77 76 78 274 275 274 78 273 272 271 38 270 269 268 64 285 40 286 287 286 40 41 42 43 44 45 46 47 48 47 46 45 44 43 42 41 40 39 38 37 36 35 34 107 106 105 104 103 102 100 101 100 102 103 104 105 106 107 34 145 144 143 142 141 139 140 139 141 142 143 144 145 34 33 32 31 30 29 28 27 26 25 4 3 2 0 1 0 2 3 4 5 194 5 193 192 193 5 6 10 152 62 61 62 185 14 184 284 283 241 93 282 281 280 279 278 276 277 276 278 279 280 281 282 93 128 95 96 127 126 125 116 124 123 122 123 124 116 115 114 113 112 7 111 110 109 223 224 225 226 4 227 228 229 230 231 232 231 230 229 228 227 4 117 118 119 120 121 301 302 301 121 120 119 195 196 197 198 199 200 201 202 203 204 205 206)

route = File.open("./complete_nodes.txt", "r").read.split(" ").map &:to_i

class Node
  attr_reader :n_id, :name
  @@node_max = 0
  def initialize(id, name)
    @n_id = id
    @name = name
  end
  def to_s
    "#{@n_id},#{@name}"
  end
end

class Edge
  attr_accessor :e_id, :depart_id, :arrival_id, :duration, :lines
  @@edge_max = 0
  @@edges = {}
  def initialize(depart_id, arrival_id, duration, line)
    @e_id = @@edge_max; @@edge_max += 1
    @depart_id= depart_id; @arrival_id = arrival_id
    @duration = duration; @lines = [*line.split(' ')]
  end
  def self.add(e)
    ex = @@edges[gen_key(e.depart_id, e.arrival_id)]
    if ex.nil?
      @@edges[gen_key(e.depart_id, e.arrival_id)] = e
    else
      e.lines[0].split(' ').each { |line| ex.lines << line }
    end
  end
  def self.gen_key(depart_id, arrival_id)
    "#{depart_id};#{arrival_id}"
  end
  def self.find(depart_id, arrival_id)
    @@edges[gen_key(depart_id, arrival_id)]
  end
  def self.all
    @@edges
  end
  def to_s
    "#{@depart_id},#{@arrival_id},#{@duration},#{@lines}"
  end
end


nodes = File.open("./nodes.csv", "r").read.split("\n").map { |line| Node.new(*line.split(",")) }
File.open("./edges.csv","r").read.split("\n").each_with_index { |line,i| Edge.add(Edge.new(*line.split(","))) if i>0 }

n_last, n_last_2, line_last = nil; duration = 0; dur_tot = 0
switch_count = 0
route.each_with_index do |n_id, i|
  n = nodes[n_id]
  e = Edge.find(n_last.n_id, n.n_id) unless n_last.nil?
  if n_last.nil?
    puts "Start from #{n.name}"
  elsif e.nil?
    puts "---> Jump made. #{n_last.name}(#{n_last.n_id}) -> #{n.name} (#{n.n_id})"
  elsif !n_last_2.nil? && n_last_2.n_id == n.n_id
    puts "After #{duration}min, turn around at #{n_last.name} and head back towards #{n.name}"
    switch_count += 1
    duration = 0
  elsif (!e.nil? && line_last.nil?)
    puts "Start on line #{e.lines.join('/')} at #{n_last.name}, towards #{n.name}"
    line_last = e.lines[0]
  elsif !e.lines.include?(line_last)
    puts "After #{duration}min, switch from line #{line_last} to line #{e.lines.join('/')} at #{n_last.name}, towards #{n.name}"
    switch_count += 1
    duration = 0
    line_last = e.lines[0]
  elsif i == route.length-1
    puts "Finish at #{n.name}"
  end
  
  unless e.nil?
    duration += e.duration.to_i 
    dur_tot += e.duration.to_i  unless e.nil?
   end
   n_last_2 = n_last; n_last = n
end
puts "Total time: #{dur_tot}min, number of switches: #{switch_count}"
