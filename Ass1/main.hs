
import RainbowAssign
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
height = 1000           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table


convertToBase :: Int -> String
convertToBase a = [toLetter (mod a nLetters)] ++ convertToBase(div a nLetters)

pwReduce :: Int -> String
pwReduce a = reverse(take pwLength (convertToBase a))