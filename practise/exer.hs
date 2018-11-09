det a b c = b^2 - 4*a*c
quadsol1 a b c = (-b - sqrt (det a b c))/2*a
quadsol2 a b c = (-b + sqrt (det a b c))/2*a

third_a x = x !! 2


hailstone :: Integer -> Integer
hailstone a 
 | mod a 2 == 0 = a `div` 2
 | otherwise = 3*a +1


-- exercise 2	

-- hailLen 1 = 0
-- hailLen a = 1 + hailLen (hailstone a)

divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]


primes :: Int -> [Int]
primes a = [i | i <- [1..a], divisors i == [] ]

-- join :: [Char] -> [[Char]]-> [Char]
-- join a [xs] = xs
-- join a (x:xs) = x ++ a ++ join a xs

pythagorean :: Int -> [(Int, Int, Int)] 
pythagorean i = [ (x, y, z) | x <- [1..i], y <- [1..i], z<- [1..i], (x^2 + y^2) == z^2 , y > x, z > y ]


-- Exercise 3
merge :: Ord a  => [a]-> [a] -> [a]
merge [] [] = []
merge [] a = a
merge a [] = a
merge (x:xs) (x2:x2s)
 | (x2 > x) = x : merge xs (x2:x2s)
 | otherwise = x2 : merge (x:xs) x2s

hailLen :: Integer -> Integer
hailLen n = hailTail 0 n
 where
 hailTail a 1 = a
 hailTail a n = hailTail (a+1) (hailstone n)


fact :: Integer -> Integer
fact a = factrecurs 1 a
 where
 factrecurs a 1 = a
 factrecurs a b = factrecurs (a*b) (b-1)
 

-- fact' :: Integer -> Integer
-- fact' a = foldl factHelper 1 [a]
-- factHelper a_ 1 = a_
-- factHelper a_ b = factHelper (a_*b) (b-1)

fact' :: Integer -> Integer
fact' a = foldr(*) 1 [1..a]



-- daysInYear :: Integer -> [Day]
-- daysInYear y = [(fromGregorian y a b) | a <- [1..12], b <- [1..31]]
--   where jan1 = fromGregorian y 1 1
--         dec31 = fromGregorian y 12 31


insert x [] = [x]
insert x (y:ys)
 | x < y = x : y : ys
 | otherwise = y : insert x ys


insertionSort :: Ord a => [a] -> [a]
insertionSort [] = []
insertionSort [x] = [x]
insertionSort (x:xs) = insert x (insertionSort xs)

-- Exercise 4

hailSeq :: Integer -> [Integer]
hailSeq 1 = [1]
hailSeq a = [a] ++ hailSeq (hailstone a)


join :: [Char] -> [[Char]]-> [Char]
join a x = foldl addFunc "" x
 where
 addFunc "" a = a
 addFunc c d = c ++ a ++ d

mergeSort :: Ord a => [a] -> [a]
-- mergeSort [] = []
mergeSort [x] = [x]
mergeSort a = merge (mergeSort left) (mergeSort right)
  where
  left = take (length (a) `div` 2) a
  right = drop (length (a) `div` 2) a



findElt :: Integer -> [Integer] -> Maybe Int
findElt a ls = findEltHelper a ls 0
 where 
 findEltHelper a [] b = Nothing
 findEltHelper a (x:xs) b
  | a == x = Just b
  | otherwise = findEltHelper a xs (b+1)



addPair :: (Integer, Integer) -> Integer
addPair (a, b) = uncurry (+) (a, b)


withoutZeros :: [Integer] -> [Integer]
withoutZeros [] = []
withoutZeros (x:xs) 
 | x == 0 = withoutZeros xs
 | otherwise = [x] ++ withoutZeros xs

qs :: Ord a => [a] -> [a]
qs [] = []
qs (x:xs) = smaller ++ [x] ++ bigger
 where
 smaller = qs [i | i <- xs, i <= x]
 bigger = qs [i | i <- xs, i >= x ]



-- insert x [] = [x]
-- insert x (y:ys)
--  | x < y = x : y : ys
--  | otherwise = y : insert x ys


-- insertionSort :: Ord a => [a] -> [a]
-- insertionSort [] = []
-- insertionSort [x] = [x]
-- insertionSort (x:xs) = insert x (insertionSort xs)