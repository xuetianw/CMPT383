import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate

-- Merging
-- why Ord a Char a does not work
-- need to ask professor for slotions

merge :: (Ord a) => [a] -> [a] -> [a]
merge [] [] = [];
merge [] x = x;
merge x [] = x;

merge (x1:x1s) (x2:x2s) -- Merging
 | (x1 <= x2) = (x1: merge (x1s)     (x2:x2s)  )
 | (x1 > x2)  = (x2: merge (x1:x1s)  (x2s)     )

--Tail Recursive Hailstone

hailstone n
 | even n = div n 2
 | otherwise = 3*n + 1

hailLen n = hailTail 0 n
  where
    hailTail a 1 = a
    hailTail a n = hailTail (a+1) (hailstone n)
-- hailTail a n = a + 1 + hailTail n

-- hailLen n = hailTail 1 n
--   where
--     hailTail a 1 = 0
--     hailTail a n = hailTail a+1 n-1

-- Factorial
--fact:: (Ord a) => a -> a
-- need to ask professor

fact n 
 | (n == 0)   = 1
 | (n > 0)    = n* fact (n-1)




fact' x = foldl(*) 1 [1..x]

daysInYear :: Integer -> [Day]
daysInYear y = [(fromGregorian y a b) | a <- [1..12], b <- [1..31]]
  where jan1 = fromGregorian y 1 1
        dec31 = fromGregorian y 12 31

isFriday :: Day -> Bool
isFriday a
 | (b == 5)    = True
 | otherwise     = False
  where (c, b) = mondayStartWeek(a)




divisors :: Int -> [Int]
divisors n = [ i | i <- [2..(n `div` 2)], n `mod` i == 0]



isPrimeDay :: Day -> Bool
isPrimeDay x
 | (divisors c == [])  = True
 |  otherwise          = False
  where (a,b,c) = toGregorian(x)


primeFridays:: Integer -> [Day]
primeFridays a = [i | i <- c, isFriday i, isPrimeDay i]
  where c = daysInYear a