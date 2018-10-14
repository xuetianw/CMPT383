------------------------------------------------------------
--
-- ColourRGBRecord.hs
-- Code sample accompanying topic 1.3.3 "Algebraic data types"
-- See README.md for details
--
-- Fundamentals of Practical Haskell Programming
-- By Richard Cook
--
------------------------------------------------------------

module ColourRGBRecord (Colour) where

data Colour = RGB
    { red :: Int
    , green :: Int
    , blue :: Int
    } deriving Show
