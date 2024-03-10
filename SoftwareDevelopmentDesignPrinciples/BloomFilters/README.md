# Bloom Filter Data Structure
A bloom filter is a set-like data structure that is more space-efficient compared to traditional set-like data structures such as hash tables or trees. The key difference is that it is probabilistic, meaning it cannot give a 100% certain answer to certain queries. More specifically, a bloom filter can tell you with 100% certainty that something is not in the set. But it cannot tell you with 100% certainty that something is in the set.

Bloom filters are typically used in OLAP databases such as Snowflake warehouse as part of the data skipping/pruning (for reading relevant underlying parquet or Snowflake specific files residing on S3) optimization scheme for a given query – use bloom filter data structure as metadata to decide which files are relevant for a given query, much smaller than indices and more load-friendly.

Similarly, LSM-tree based key-value (KV) stores suffer from severe read amplification because searching a key requires to check multiple SSTables. To reduce extra I/Os, Bloom filters are usually deployed in KV stores to improve read performance.

A bloom filter data structure is actually an array of bits. Each item you pass to your bloom filter gets mapped to a location in the bloom filter, by hashing the incoming item. When you later query your bloom filter, you hash the input and turn the hash into a location inside the bitfield. Next you check if the corresponding bits are set to 1 or 0. If they have been set to 1, than the item is in the set.

The hashing of the input is the reason why you can be certain of absence, but not certain of presence. If a bit is 0 you’re certain that nothing ever ended up on that index. But if it’s 1, it could be due to a ‘hash collision’ – simply because multiple hashes mapped to the same position.

Essentially, a good bloom filter gives a small number of false positives possible, as this avoids the (expensive) lookup in the real data set. One way of mitigating the false positives is by adding more hash functions, thus decreasing our odds of multiple items mapping to the same location in our bitfield.

See the companion [code example](/SoftwareDevelopmentDesignPrinciples/BloomFilters). In this example code, a very simplistic Bloom Filter maps the input string to a location field in our bitfield. When we store data in it using f.Set("www.yahoo.com") we’ll flip a bit to 1 to indicate that it has been set. When we run f.Get("www.yahoo.com") we use the same hash to find out if the bit has been set. That is all there is to create a basic bloom filter.