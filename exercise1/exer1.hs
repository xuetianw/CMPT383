det a b c = b^2 - 4*a*c
quadsol1 a b c = (-b - sqrt (det a b c))/2*a
quadsol2 a b c = (-b + sqrt (det a b c))/2*a


third_a lst = lst!!2

third_b (_:_:x:_) = x


 --hailstone a 

find_mod n = mod n 2

hailstone n
 | even n = div n 2
 | otherwise = 3*n + 1
