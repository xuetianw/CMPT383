-- ghc -O2 -with-rtsopts="-N8" -threaded concurrent1.hs
-- time ./concurrent1

import Control.Parallel
import Control.Parallel.Strategies

n = 10000000000

myPowerSeq :: Int -> Int -> Int -> Int
myPowerSeq a _ 0 = a
myPowerSeq a x y = seq newacc (myPowerSeq newacc x (y-1))
    where newacc = a*x

-- some function to take a reasonable amount of time in examples...
calculation i = myPowerSeq i 1 n
slowCalc i = myPowerSeq i 1 (n `div` 50)
fastCalc i = myPowerSeq i 1 (n `div` 500000)

calcA = a + b
    where a = calculation 1
          b = calculation 2


calcB = (a `par` b) `pseq` (a + b)
    where a = calculation 1
          b = calculation 2


calcC = map slowCalc [0..100]


calcD = parMap rseq slowCalc [0..100]


calcE = map fastCalc [0..1000000]
calcF = parMap rseq fastCalc [0..1000000]





main :: IO ()
main = print $ calcE