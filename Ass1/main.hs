
import RainbowAssign
import qualified Data.Map as Map

-- main :: IO ()
-- main = do
--   generateTable
--   res <- test2 10000
--   print res
pwLength, nLetters, width, height :: Int
filename :: FilePath
pwLength = 8            -- length of each password
nLetters = 5            -- number of letters to use in passwords: 5 -> a-e
width = 40              -- length of each chain in the table
height = 10           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table


pwReduce :: Hash -> Passwd
pwReduce hash = map toLetter (reverse (take pwLength (helper (fromEnum hash))))
    where
        helper :: Int -> [Int]
        helper hash = [hash `mod` nLetters] ++ helper (hash `div` nLetters)


rainbowTable :: Int -> [Passwd] -> Map.Map Hash Passwd
rainbowTable n passwords = Map.fromList (listToTuple (rainbowTableRecur n passwords) passwords)
    where
        rainbowTableRecur :: Int -> [Passwd] -> [Hash]
        rainbowTableRecur 0 passwords = getHashes passwords
        rainbowTableRecur n passwords = rainbowTableRecur (n-1) (map pwReduce (getHashes passwords))


getHashes :: [Passwd] -> [Hash]
getHashes passwords = map pwHash passwords

listToTuple :: [Hash] -> [Passwd] -> [(Hash, Passwd)]
listToTuple [] [] = []
listToTuple (x1:x1s) (x2:x2s) =  [(x1,x2)] ++ listToTuple x1s x2s

-- rainbowTable tableWidth [passwords] = rainbowTable (tableWidth-1) [pwReduce (pwHash passwords)] 
-- = Map.fromList[(1726491528,"abcdeabc")]

-- findStartPw :: Map.Map Hash Passwd -> Int -> Hash -> Maybe Passwd
-- findStartPw table (-1) hash = Nothing
-- findStartPw table width hash
--     | result == Nothing = findStartPw table (width-1) (pwHash (pwReduce hash))
--     | otherwise         = result
--     where result = Map.lookup hash table

-- findRealPw :: Maybe Passwd -> Hash -> Maybe Passwd
-- findRealPw (Just startPw) hash
--     | pwHash startPw == hash = Just startPw
--     | otherwise = findRealPw (Just (pwReduce (pwHash startPw))) hash

-- findPassword :: Map.Map Hash Passwd -> Int -> Hash -> Maybe Passwd
-- findPassword table width hash
--     | startPw == Nothing = Nothing
--     | otherwise = findRealPw startPw hash
--     where startPw = findStartPw table width hash

findStartPw :: Map.Map Hash Passwd-> Int -> Hash -> Maybe Passwd
findStartPw table (-1) hash = Nothing
findStartPw table width hash 
 | Map.lookup hash table == Nothing = findStartPw table (width-1) (pwHash (pwReduce hash))
 | otherwise = Map.lookup hash table 

findRealPw :: Maybe Passwd -> Hash -> Maybe Passwd
findRealPw (Just password) hash 
 | (pwHash password) == hash = Just password
 | otherwise = findRealPw (Just (pwReduce (pwHash password)) ) hash


findPassword :: Map.Map Hash Passwd-> Int -> Hash -> Maybe Passwd
findPassword table width hash 
 | findStartPw table width hash == Nothing = Nothing 
 | otherwise = findRealPw (findStartPw table width hash) hash

generateTable :: IO ()
generateTable = do
  table <- buildTable rainbowTable nLetters pwLength width height
  writeTable table filename
