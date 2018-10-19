-- Rational Numbers
import Data.Ratio

rationalSum :: (Integral a) => a -> [Ratio a]
rationalSum a = [ (x % y) | x <- [1..a], y <- [1..a], y == (a - x)]

-- Lowest Terms Only

rationalSumLowest :: (Integral a) => a -> [Ratio a]
rationalSumLowest a = [ (x % y) | x <- [1..a], y <- [1..a], y == (a - x), gcd x y ==1]
