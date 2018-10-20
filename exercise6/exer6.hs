-- Rational Numbers
import Data.Ratio
import Data.List.Split

rationalSum :: (Integral a) => a -> [Ratio a]
rationalSum a = [ (x % y) | x <- [1..a], y <- [1..a], y == (a - x)]

-- Lowest Terms Only

rationalSumLowest :: (Integral a) => a -> [Ratio a]
rationalSumLowest a = [ (x % y) | x <- [1..a], y <- [1..a], y == (a - x), gcd x y ==1]

-- All Rational Numbers
rationals :: (Integral a) => [Ratio a]
rationals = rationalSumLowestRec 0
 where 
  rationalSumLowestRec a = (rationalSumLowest a) ++ (rationalSumLowestRec (a+1))


-- Input/Output
splitFunc :: [Char] -> [[Char]]
splitFunc line_ = splitOn "\n" line_

removeEmptyNumer :: [[Char]] -> [[Char]]
removeEmptyNumer [] = []
removeEmptyNumer (x:xs) 
 | x == "" = removeEmptyNumer xs
 | otherwise = [x] ++ removeEmptyNumer xs


f :: [String] -> [Int]
f = map read

sumFile :: IO ()
sumFile = do
    line <- readFile "input.txt"
    let a = splitFunc line
    let b = removeEmptyNumer a
    let c = f b
    print (sum c)