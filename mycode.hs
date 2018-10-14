import System.Random

half x = x/2
isBig a = a > 100
listify x y = [x, y]


-- Pattern Matching
myAnd  False _ = False
myAnd True a = a


isZero 0 = True
isZero _ = False

myLength [] = 0
myLength (_:xs) = 1 + myLength xs

myLength' (xs) = myLengthtail 0 (xs)
  where myLengthtail a [] = 0
        myLengthtail a (_:xs) = 1 + myLengthtail a (xs)

listEqual [] (_:_) = False
listEqual (_:_) [] = False
listEqual [] [] = True
listEqual (x:xs) (y:ys) = x == y && listEqual xs ys
-- listEqual _ _ = False  --and use the first-match-only behaviour.)


hasEvenLength [] = True     -- base case: length 0 is even
hasEvenLength [_] = False   -- base case: length 1 is odd
hasEvenLength (_:_:rest) = hasEvenLength rest


-- Conditionals
mySignum x
  | x>0         = 1
  | x<0         = -1
  | otherwise   = 0


word n = case n of
  1 -> "one"
  2 -> "two"
  3 -> "three"
  _ -> "???"


describeList lst = "The list is " ++ case lst of
  _:_:_:_:_  -> "fairly long"     --  >= 4 elements
  _:_        -> "short"           --  >= 1 element
  []         -> "empty"
  
  
add :: Int -> Int -> Int
add a b = a + b

type Count = Int
processString ::  [Char] -> Count
--processString ::  [Char] -> Int -- also work
processString list = 2

addPair :: Num a=> (a, a) -> a
addPair = uncurry (+)
-- addPair (2,3)
 

--List Comprehensions
-- qs is really important to understand
-- try insertion sort

qs [] = []
qs (x:xs) = smaller ++ [x] ++ larger
  where smaller = qs [a | a<-xs, a<=x]
        larger = qs [a | a<-xs, a>x]

-- qs [3,2,1,4,5] = qs [a | a<-[2,1,4,5], a<=3 ] ++ [3] ++  qs [a | a<-[2,1,4,5], a>3 ]

-- qs [a | a<-[2,1,4,5], a<=3 ] = qs [2,1] = [1] ++ [2]
-- qs [a | a<-[2,1,4,5], a>3 ] = qs [4,5] = [4] ++ [5]


absolute x
  | x > 0     = x
  | x < 0     = -x
  | x == 0    = x



somePowers x = [x, sq x, sq (sq x)]
  where sq n = n*n




-- Recursion

-- this is a normal recursive function
myPower _ 0 = 1
myPower x y = x * myPower x (y-1)



myPower' x y
  | y==0      = 1
  | even y    = half*half
  | odd y     = x*half*half
  where half = myPower' x (div y 2)




myPowerTailRec' x y = myPowerTailRec 1 x y
    where
    myPowerTailRec a _ 0 = a
    myPowerTailRec a x y = myPowerTailRec (x*a) x (y-1)


--another wat if implementing Tailrecursion
-- myPowerTailRec' x y = 
--     let myPowerTailRec a _ 0 = a
--         myPowerTailRec a x y = myPowerTailRec (x*a) x (y-1)
--     in myPowerTailRec 1 x y
    

factorial 0 = 1
factorial n = n * factorial (n-1)


-- factorialTailRec a 0 = a
-- factorialTailRec a n = factorialTailRec (n*a) (n-1)
-- factorial' n = factorialTailRec 

factorial' n = factorialTailRec 1 n
    where
    factorialTailRec a 0 = a
    factorialTailRec a n = factorialTailRec (n*a) (n-1)


hailstone n
 | even n    = div n 2
 | otherwise = 3*n + 1


divisors n = filter (divides n) [2..(n `div` 2)]
  where divides a b = (a `mod` b) == 0

-- wrote a random meaningless function for testing
mySum2 n = mySum2tail 0 n
 where mySum2tail a 1 = a + 1
       mySum2tail a c = mySum2tail (a+c) (c-1)

--wrote a recursive call
mySum2' 0 = 0
mySum2' n = n + mySum2' (n-1)

mySum [] = 0
mySum (x:xs) = x + mySum xs

mySum' xs = foldl (+) 0 xs
mySum'[1,2,3] = foldl (+) ((+) ((+) 0 1) 2) [3]

lst :: [Int]
lst = [1,2,3]

lstlst :: [[Int]]
lstlst = [[1,2,3], [4,5]]

myReverse xs = foldl prefix [] xs
    where prefix xs x = x:xs



myReverse' xs = foldr postfix [] xs
    where postfix x xs = x:xs


join :: String -> ([String] -> String)
join _ [] = ""
join _ [a]= a
join a (x:xs) = x ++ a ++ join a xs




commajoin :: [String] -> String
commajoin = join ", "

myConcat :: [[a]] -> [a]
myConcat xs = foldl (++) [] xs


-- why this is working? 
-- ask TA or professor!!
myConcat' :: [[a]] -> [a]
myConcat' = foldl (++) []


myDiv :: (Int, Int) -> Int
myDiv (n,d) = div n d


addPairwise :: Num a => [a] -> [a] -> [a]
addPairwise xs ys = map (uncurry (+)) (zip xs ys)


reverseJoin :: String -> [String] -> String
reverseJoin s = (join s) . reverse


joinPrimes :: String -> String
joinPrimes = (flip join) ["2", "3", "5", "7", "11"]


myReverse'' :: [a] -> [a]
myReverse'' xs = foldl (flip (:)) [] xs


half_of' :: Float -> Float
half_of' = \x -> x/2

-- (foldl (+)) 0 [1,2,3,4,5] 
-- == (+) ((+) ((+) ((+) ((+) 0 1) 2) 3) 4) 5


--foldl (flip (:)) [] "testing" 
-- == flip (:) (flip (:) (flip (:)(flip (:) (flip (:) (flip (:) (flip (:) [] "t") "e") "s") "t") "i" ) "n" ) "g"


addToEach :: Num a => a -> [a] -> [a]
addToEach n lst = map (\x -> x+n) lst

-- why it is not this? foldl (+) ((+)0 1) [2,3] 
-- need to ask TA or professor

-- foldl (+) 0 [1,2,3]
-- == foldl (+) (0+1) [2,3]
-- == foldl (+) ((0+1)+2) [3]
-- == foldl (+) (((0+1)+2)+3) []
-- == ((0+1)+2)+3
-- == (1+2)+3
-- == 3+3
-- == 6

--foldl (+) ((+) 0 1) [2,3]
-- import Data.List
-- :t foldl'   cannot understand the type need to ask professor

myPowerTailRec :: Int -> Int -> Int -> Int
myPowerTailRec a _ 0 = a
myPowerTailRec a x y = myPowerTailRec (x*a) x (y-1)

myPowerTailStrict a _ 0 = a
myPowerTailStrict a x y = (myPowerTailStrict $! (a*x)) x (y-1)


myPowerSeq a _ 0 = a
myPowerSeq a x y = seq newacc (myPowerSeq newacc x (y-1))
    where newacc = a*x

-- myPowerSeq 1 2 3 ==
--  myPowerSeq 2 2 2 == myPowerSeq 4 2 2 == myPowerSeq 8 2 1

funnyDivisors n   = map pred (divisors (n*2))
funnyDivisors' n  = map pred $ divisors $ n*2
funnyDivisors'' n = (map pred) . divisors . (*2) $ n
funnyDivisors'''  = (map pred) . divisors . (*2)

-- funnyDivisors1 = (map pred) . divisors . (*2) $ n


-- mySum( myReverse [1,2,3,4,5] )
-- mySum.myReverse [1,2,3,4,5]



-- calcB = (a `par` b) `pseq` (a + b)
--     where a = calculation 1
--           b = calculation 2

-- euqal function
-- calcB = pseq (par a  b)  (a + b)
--     where a = calculation 1
--           b = calculation 2

-- Int has space limit
-- let base=2 :: Int
-- let power=64 :: Int
-- base ^ power


secondLine :: IO String
secondLine = do
    line1 <- getLine
    line2 <- getLine
    return line2

secondLine' :: IO String
secondLine' =
  getLine >>= (\line1 -> getLine >>= (\line2 -> return line2))

secondLine'' :: IO String
secondLine'' = do
    line1 <- getLine
    getLine

succInput :: IO ()
succInput = do
    text <- getLine
    let succtext = map succ text
    putStrLn succtext


succInput' :: IO ()
succInput' = do
    text <- getLine
    succtext <- return $ map succ text
    putStrLn succtext

-- convert a list of values into a histogram
histogram :: (Enum a, Eq a, Ord a) => [a] -> [String]
histogram vals = bars
    where
    counts = [length $ filter (==i) vals
             | i <- [(minimum vals)..(maximum vals)]]
    bars = [take n $ repeat 'X' | n <- counts]


threeRand :: [Int]
threeRand = 
    let gen0 = mkStdGen 1234 -- gen0 :: StdGen
        (rand0, gen1) = randomR (1, 100) gen0
        (rand1, gen2) = randomR (1, 100) gen1
        (rand2, _)    = randomR (1, 100) gen2
    in [rand0, rand1, rand2]


threeRand' :: IO [Int]
threeRand' = do
    gen0 <- newStdGen
    let
        (rand0, gen1) = randomR (1, 100) gen0
        (rand1, gen2) = randomR (1, 100) gen1
        (rand2, _)    = randomR (1, 100) gen2
    return [rand0, rand1, rand2]



threeRand'' :: IO [Int]
threeRand'' = do
    gen0 <- newStdGen
    return $ take 3 $ randomRs (1, 100) gen0



-- generate n random integers
randInts :: Int -> Int -> Int -> IO [Int]
randInts n minval maxval = do
    gen <- newStdGen
    return $ take n $ randomRs (minval, maxval) gen

printHisto :: IO ()
printHisto = do
    vals <- randInts 1000 1 20
    let bars = histogram vals
    mapM_ putStrLn bars

    