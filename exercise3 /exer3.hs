import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate


-- Merging
-- instance Enum Int -- Defined in ‘GHC.Enum’
merge :: (Ord a) => [a] -> [a] -> [a]
merge [] x = x
merge x [] = x
merge (x1:x1s) (x2:x2s)
 | x1 < x2  = [x1] ++ merge x1s (x2:x2s)
 | x1 >= x2  = [x2] ++ merge (x1:x1s) x2s



hailstone n
 | even n = div n 2
 | otherwise = 3*n + 1

-- hailLen 1 = 0
-- hailLen n = 1 + hailLen (hailstone n)

-- Tail Recursive Hailstone
hailLen :: Int -> Int
hailLen n = hailTail 0 n
  where
    hailTail a 1 = a 
    hailTail a n = hailTail (a+1) (hailstone n)


-- Factorial

fact :: Integer -> Integer
fact n = factrecurs n
 where
  factrecurs 1 = 1
  factrecurs n = n * factrecurs(n-1)


-- fact' :: Integer -> Integer
-- fact' n 
--  | n == 1  = 1
--  | n > 1   = n * fact'(n-1)


fact' :: Integer -> Integer
fact' a = foldr(*) 1 [1..a]


 -- Haskell Library and Dates
daysInYear :: Integer -> [Day]
daysInYear y = [jan1..dec31]
  where jan1 = fromGregorian y 1 1
        dec31 = fromGregorian y 12 30

getDay (y,m,d) = d
getDay2 (m,d) = d

isFriday :: Day -> Bool
isFriday a 
 | getDay2 (mondayStartWeek a) == 5    = True
 | otherwise         				   = False


divisors :: Int -> [Int]
divisors n = [ i | i <- [2..(n `div` 2)], n `mod` i == 0]


-- reference link: https://two-wrongs.com/haskell-time-library-tutorial
isPrimeDay :: Day -> Bool
isPrimeDay a 
 | divisors (getDay (toGregorian a)) == []  = True
 | otherwise 								= False

primeFridays :: Integer -> [Day]
primeFridays a = [ i | i <- (daysInYear a), isPrimeDay i, isFriday i]