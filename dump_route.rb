# Expects a list of connected nodes in routes (connections should be in edges.csv).
# For any discontinuous jumps, the jump is printed and the calculation is picked up at next location.
route = %w(206 205 204 203 202 201 200 199 198 197 196 195 194 193 192 191 190 189 188 187 186 185 184 183 182 181 180 179 178 177 247 246 245 244 243 242 241 240 239 238 248 249 237 236 235 234 233 138 137 136 135 134 133 132 131 130 129 128 127 126 99 98 97 96 95 94 93 92 91 90 89 88 87 86 85 84 83 82 81 80 79 78 77 76 107 106 105 104 103 102 101 100 139 140 141 142 143 144 145 146 147 108 109 110 111 112 113 114 115 116 125 124 123 122 121 120 119 118 117 232 231 230 229 228 227 226 225 224 223 176 74 73 52 53 54 56 255 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 222 148 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168 169 170 171 173 174 175 172 51 50 49 75 250 251 252 253 254 55 275 274 273 272 271 270 269 268 267 266 265 264 263 262 261 260 259 258 257 256 276 277 278 279 280 281 282 283 284 285 286 287 288 289 290 291 292 293 294 295 296 297 298 299 300 48 47 46 45 44 43 42 41 40 39 38 37 36 35 34 33 32 31 30 29 28 27 26 25 4 3 2 1 0 5 301 302 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 221 220 219 218 217 216 215 214 213 212 211 210 209 207 208).map &:to_i

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
  attr_reader :e_id, :depart_id, :arrival_id, :duration, :line
  @@edge_max = 0
  @@edges = {}
  def initialize(depart_id, arrival_id, duration, line)
    @e_id = @@edge_max; @@edge_max += 1
    @depart_id= depart_id; @arrival_id = arrival_id
    @duration = duration; @line = line
  end
  def self.add(e)
    @@edges[gen_key(e.depart_id, e.arrival_id)] = e
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
    "#{@depart_id},#{@arrival_id},#{@duration},#{@line}"
  end
end


nodes = File.open("./nodes.csv", "r").read.split("\n").map { |line| Node.new(*line.split(",")) }
File.open("./edges.csv","r").read.split("\n").map { |line| Edge.add(Edge.new(*line.split(","))) }

n_last = nil; e_last = nil
route.each_with_index do |n_id, i|
  n = nodes[n_id]
  e = Edge.find(n_last.n_id, n.n_id) unless n_last.nil?
  #puts "#{n} -> #{e}"
  if n_last.nil?
    puts "Start from #{n.name}"
  elsif e.nil?
    puts "Jump made. #{n_last.name} -> #{n.name}"
  elsif (!e.nil? && e_last.nil?) || e_last.line != e.line
    puts "At #{n_last.name}, travel on line #{e.line} towards #{n.name}"
  elsif i == route.length-1
    puts "Finish at #{n.name}"
  end
  n_last = n; e_last = e unless e.nil?
end