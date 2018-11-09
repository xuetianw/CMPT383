-- Hailstone, Again
hailstone :: (Num a, Integral a) => a -> a
hailstone n
 | even n = div n 2
 | otherwise = 3*n + 1

-- hailSeq 1 = [1]
-- hailSeq a = [a] ++ hailSeq (hailstone a)
hailSeq :: Integral a => a -> [a]
hailSeq x 
  | x == 1 = [1]
  | otherwise = [x] ++ hailSeq (hailstone x)


-- hailLen :: Integer -> Integer
-- hailLen 1 = 0
-- hailLen a = 1 + hailLen (hailstone a)

hailLen :: (Num a, Integral a) => a -> a
hailLen n = hailTail 0 n
 where
 	hailTail a 1 = a
 	hailTail a n = hailTail (a+1) (hailstone n)


hailSeq' :: Int -> [Int]
hailSeq' a = mytakeWhile (/=1) (myIterate hailstone a) ++ [1]

myIterate :: (a -> a) -> a -> [a]
myIterate f x = [x] ++ myIterate f (f x)


mytakeWhile :: (a -> Bool) -> [a] -> [a]
mytakeWhile f [] = []
mytakeWhile f (x:xs) 
 | f x =  [x] ++ mytakeWhile f xs
 | otherwise = []


-- hailSeq :: Integer -> [Integer]
-- hailSeq n = hailSeqRecur n
--  where 
--  	hailSeqRecur 1 = [1]
--  	hailSeqRecur n = [n] ++ hailSeqRecur (hailstone n)




-- -- Non-recursive
-- hailLen 1 = 0
-- hailLen n = 1 + hailLen (hailstone n)

-- hailSeq' :: Integer -> [Integer]
-- -- hailSeq' n = take (hailLen n + 1) (iterate hailstone n)
-- hailSeq' n = (myTakeWhile (/=1) (iterate hailstone n)) ++ [1]

-- myTakeWhile :: (a -> Bool) -> [a] -> [a]
-- myTakeWhile f [] = []
-- myTakeWhile f (x:xs) 
--  | f x         = [x] ++ myTakeWhile f (xs)
--  | otherwise   = []


-- myIterate :: (a -> a) -> a -> [a]
-- myIterate f n = [n] ++ myIterate f (f n)
-- hailSeq'' n = take (hailLen n + 1) (myIterate hailstone n)

-- Joining Strings, Again

join :: String -> [String] -> String
join a xs = foldl addFunc' "" xs
 where
 	addFunc' :: String -> String -> String
 	addFunc' b x2 
 		|b == ""	= x2
 		|otherwise 	= b ++ a ++ x2

myFoldl f a [] = a
myFoldl f a (x:xs) = myFoldl f (f a x) xs





	 -- b ++ ((++) a x)
-- join sep xs = foldr add "" xs
--          where add a b
--                   | b == ""    = a ++ b
--                   | otherwise  = a ++ sep ++ b

--Merge Sort

merge :: Ord a => [a] -> [a] -> [a]
merge  a [] = a
merge  [] a = a
merge (x1:x1s) (x2:x2s) 
	| x1 < x2	= x1 : merge x1s (x2:x2s)
	| x2 <= x1  = x2 : merge (x1:x1s) x2s

mergeSort :: Ord a => [a] -> [a]
mergeSort [] = []
mergeSort [x] = [x]
mergeSort lst = merge (mergeSort left) (mergeSort right)
	where
		left = take (length lst `div` 2) lst
		right = drop (length lst `div` 2) lst

-- found: return idx
-- not found: return -1

findEltHelper :: Eq a => a -> [a] -> Int -> Int
findEltHelper k [] _ = -1
findEltHelper k (x:xs) idx
	| k == x = idx
	| otherwise = findEltHelper k xs (idx + 1)

findElt :: Eq a => a -> [a] -> Maybe Int
findElt k lst 
	| result == -1 = Nothing
	| otherwise = Just result
	where result = findEltHelper k lst 0

-- Maybe:
-- Implement function called cvtMaybe
-- cvtMaybe (Just 3) ===> 3
-- cvtMaybe Nothing ===> error
-- cvtMaybe (Just 9.6) ==> 9.6
-- cvtMaybe (Just True) ==> True
-- cvtMaybe (Just []) ==> []
-- (10 pts)

-- 5 pts for type signature
-- 5 pts for correctly implementation

cvtMaybe :: Maybe a -> a
cvtMaybe (Just a) = a

-- Maybe:
-- cntMaybe
-- cntMaybe [] ===> 0
-- cntMaybe [Just 3, Just 2, Nothing] ===> 2
-- cntMaybe [Nothing, Nothing] ===> 0
-- cntMaybe [Just 6.8, Just 9.8] ==> 2


cntMaybe :: [Maybe a] -> Int
cntMaybe lst = cntMaybeRecur lst 0
	where
		cntMaybeRecur [] count = count
		cntMaybeRecur (Nothing:xs) count = cntMaybeRecur xs count
		cntMaybeRecur (_:xs) count = cntMaybeRecur xs (count + 1)


-- Maybe:
-- rmNothing
-- rmNothing [] ====> []
-- rmNothing [Nothing, Nothing, Just 3] ====> [Just 3]
-- rmNothing [Just 5, Just 8] ===> [Just 5, Just 8]
-- rmNothing [Just 7.2, Nothing, Just 2.3] ===> [Just 7.2, Just 2.3]

rmNothing :: [Maybe a] -> [Maybe a]

rmNothing ls = rmNothingRec ls []
	where
		rmNothingRec [] a = a
		rmNothingRec (Nothing:xs) a = rmNothingRec xs a
		rmNothingRec (x:xs) a = rmNothingRec xs (x:a)



-- cntMaybe (x:xs) = cntMaybeRecur (x:xs) 0
-- 	where 
-- 		cntMaybeRecur [] a  = a
-- 		cntMaybeRecur (x:xs) a
-- 			| x == Nothing   = cntMaybeRecur xs a
-- 			| otherwise       = cntMaybeRecur xs (a + 1)


