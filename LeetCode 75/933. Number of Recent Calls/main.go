package main

func main() {

}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

type RecentCounter struct {
	requests []int
	cnt      int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.requests) > 0 && this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
		this.cnt--
	}

	this.cnt++
	this.requests = append(this.requests, t)

	return this.cnt
}
