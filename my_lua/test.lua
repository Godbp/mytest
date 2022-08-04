local k1 = KEYS[1]
local field = ARGV[1]
local val = ARGV[2]
return redis.call('hset', k1, field, val)