------------------------------------------------------------
--
-- Program.hs
-- Code sample accompanying topic 1.2.3 "GHCi and interactive Haskell"
-- See README.md for details
--
-- Fundamentals of Practical Haskell Programming
-- By Richard Cook
--
------------------------------------------------------------

module Main (main, prettyPrint) where

type Port = Int

data Address = Address Int Int Int Int Port

prettyPrint :: Address -> IO ()
prettyPrint (Address ip0 ip1 ip2 ip3 port)
    = putStrLn $
        show ip0 ++ "." ++
        show ip1 ++ "." ++
        show ip2 ++ "." ++
        show ip3 ++ ":" ++ show port

main :: IO ()
main = prettyPrint (Address 127 0 0 1 80)
