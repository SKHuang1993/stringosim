package stringosim

import (
    "strings"
)

type CosineSimilarityOptions struct {
    CaseInsensitive bool
    NGramLength     int
}

var DefaultCosineSimilarityOptions = CosineSimilarityOptions{
    CaseInsensitive: false,
    NGramLength:     2,
}

func Cosine(s []rune, t []rune, options ...CosineSimilarityOptions) float64 {
    if len(s) == 0 {
        if len(t) == 0 {
            return float64(0.0)
        } else {
            return float64(1.0)
        }
    } else {
        if len(t) == 0 {
            return float64(1.0)
        }
    }
    opt := DefaultCosineSimilarityOptions
    for _, option := range options {
        opt = option
        break
    }
    var sGrams, tGrams map[string]int
    if opt.CaseInsensitive {
        sGrams = GetNGram(strings.ToLower(string(s)), opt.NGramLength)
        tGrams = GetNGram(strings.ToLower(string(t)), opt.NGramLength)
    } else {
        sGrams = GetNGram(string(s), opt.NGramLength)
        tGrams = GetNGram(string(t), opt.NGramLength)
    }
    return 1.0 - float64(DotProductNGrams(sGrams, tGrams))/NormNGram(sGrams)/NormNGram(tGrams)
}
