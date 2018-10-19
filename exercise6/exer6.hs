-- Rational Numbers
import Data.Ratio
rationalSumHelper :: (Integral a) => a -> [Ratio a]
rationalSumHelper a = [ (x % y) | x <- [1..a], y <- [1..a], y == (a - x)]