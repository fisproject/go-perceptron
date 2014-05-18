perceptron
============

## Feature
-  golang clone from python (http://en.wikipedia.org/wiki/Perceptron)
- A perceptron learns to perform a binary NAND function 

## TODO
- 本当はこんなJSONをparseしたかったけどUnmarshalが上手くいかなかった

<pre>
{
    "training_set_count": 4,
    "training_set": [
        {
            "input_vector": "1,0,0",
            "desired_output": 1
        },
        {
            "input_vector": "1,0,1",
            "desired_output": 1
        },
        {
            "input_vector": "1,1,0",
            "desired_output": 1
        },
        {
            "input_vector": "1,1,1",
            "desired_output": 0
        }
    ]
}
</pre>