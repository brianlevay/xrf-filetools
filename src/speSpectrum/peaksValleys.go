package speSpectrum

import ()

func getPeaks(counts []uint64, bgMult uint64) [][]uint64 {
    var peaks [][]uint64
    var base uint64
    var delPrev, delNext uint64
    var thresh uint64
    var row []uint64
    
    base = 0
    nChannels := len(counts)
    for i := 1; i < (nChannels-1); i++ {
        delPrev = counts[i] - counts[i-1]
        delNext = counts[i+1] - counts[i]
        thresh = bgMult*base
        if (delPrev >= 0) && (delNext <= 0) {
            if (counts[i] >= thresh) {
                row = make([]uint64, 2)
                row[0] = uint64(i)
                row[1] = counts[i]
                peaks = append(peaks, row)
            }
        } else if (delPrev <= 0) && (delNext >= 0) {
            base = counts[i]
        }
    }
    return peaks
}
