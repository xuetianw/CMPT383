-- Built-In Functions

myIterate :: (a -> a) -> a -> [a]
myIterate f n = [n] ++ myIterate f (f n)

myTakeWhile :: (a -> Bool) -> [a] -> [a]
myTakeWhile f [] = []
myTakeWhile f (x:xs)
 | f x   		 = [x] ++ myTakeWhile f xs
 | otherwise 	 = []


-- hailstone n
--  | even n = div n 2
--  | otherwise = 3*n + 1

-- hailSeq' n = (myTakeWhile (/=1) (iterate hailstone n)) ++ [1]

-- Pascal's Triangle

-- helper :: [b] -> [(b, b)]
-- helper prev = zip prev (tail prev)

add :: Num a => (a, a) -> a
add (a,b) = a + b

-- tuppleSumHelper :: Num b => [(b, b)] -> [b]
-- tuppleSumHelper c = map add c

tuppleSum :: Num a => [a] -> [a]
tuppleSum prev = [1] ++ map add (zip prev (tail prev)) ++ [1]

-- why do we need tyoe Eq ?  ask professor
pascal :: (Ord a, Num a) => a -> [a]
pascal 0 = [1]
pascal 1 = [1,1]
pascal n = tuppleSum ( pascal (n - 1) )

-- pascal n = [1] ++ zip prev (tail prev) ++ [1]
--  where prev = pascal (n-1)
-- c = [(1,3),(3,3),(3,1)]



-- Pointfree Addition

addPair :: Num a=> (a, a) -> a
addPair = uncurry (+)

-- addPair' (x,y) = uncurry (+) (x, y)
-- addPair = addPair'


-- Pointfree Filtering

--my own way
withoutZeros' [] = []
withoutZeros' (x:xs) 
	| x /= 0    = x: withoutZeros' xs
	| otherwise = withoutZeros' xs


withoutZeros = withoutZeros'



-- test::Real a => a -> [a] -> [a]
-- test a b = [x | x <- b, x/=a]

-- withoutZeros:: Real a => [a] -> [a]
-- withoutZeros = test 0

-- Exploring Fibonacci
fib :: Integer -> Integer
fib 0 = 0
fib 1 = 1
fib n = fib (n-1) + fib (n-2)

fibs = map fib [0..]




things :: [Integer]
things = 0 : 1 : zipWith (+) things (tail things)
