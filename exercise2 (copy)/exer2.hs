hailstone n
 | even n = div n 2
 | otherwise = 3*n + 1

hailLen 1 = 0
hailLen n = 1 + hailLen (hailstone n)


divisors :: Int -> [Int]
divisors n = [ i | i <- [2..(n `div` 2)], n `mod` i == 0]


-- divisors n = filter (divides n) [2..(n `div` 2)]
--   where divides a b = (a `mod` b) == 0

primes :: Int -> [Int]
primes n = [i | i <- [2..n], divisors i == []]


join :: String -> [String] -> String
join _ [] = ""
join _ [a]= a
join a (x:xs) = (++) x a ++ join a xs


--test [[xs]] ++ [x]] = [x]

-- join :: String -> ([String] -> String)
-- commajoin :: [String] -> String
-- commajoin = join ", "

pythagorean :: Int -> [(Int, Int, Int)]
pythagorean n = [ (x,y,z) | x <- [1..n], y <- [1..n], z <- [1..n], x^2 + y^2 == z^2, x < y, y < z ]	