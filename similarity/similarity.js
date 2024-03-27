function calSimilarity(text1, text2) {
    var input1 = text1.match(/\b\w+\b/g)
    var input2 = text2.match(/\b\w+\b/g)
    const distance = intersectionCount(input1, input2)
    const similarity = (distance / (input1.length + input2.length - distance)).toFixed(2)
    console.log(similarity)
    return similarity
}

function intersectionCount(a, b) {
    console.log(a)
    console.log(b)
    var intersection = {}, count = 0
    for (i of a) intersection[i] = true
    for (i of b) if (intersection[i]) count++
    return count
}

function calculateSimilarity(article1, article2) {
    const words1 = article1.match(/\b\w+\b/g);
    const words2 = article2.match(/\b\w+\b/g);

    const vector1 = buildWordVector(words1);
    const vector2 = buildWordVector(words2);

    const similarity = calculateCosineSimilarity(vector1, vector2);
    return similarity;
}

function buildWordVector(words) {
    const vector = {};
    for (const word of words) {
        vector[word] = (vector[word] || 0) + 1;
    }

    return vector;
}

function calculateCosineSimilarity(vector1, vector2) {
    const dotProduct = calculateDotProduct(vector1, vector2);
    const magnitude1 = calculateMagnitude(vector1);
    const magnitude2 = calculateMagnitude(vector2);

    if (magnitude1 === 0 || magnitude2 === 0) {
        return 0;
    }

    const similarity = dotProduct / (magnitude1 * magnitude2);
    return similarity;
}

function calculateDotProduct(vector1, vector2) {
    let dotProduct = 0;

    for (const key in vector1) {
        if (vector2.hasOwnProperty(key)) {
            dotProduct += vector1[key] * vector2[key];
        }
    }

    return dotProduct;
}

function calculateMagnitude(vector) {
    let sumOfSquares = 0;

    for (const key in vector) {
        sumOfSquares += vector[key] ** 2;
    }

    const magnitude = Math.sqrt(sumOfSquares);
    return magnitude;
}

const article1 = "You should spend about 40 minutes on this task. Write about the following topic: It is important for people to take risks, both in their professional lives and their personal lives. Do you think the advantages of taking risks outweigh the disadvantages? Give reasons for your answer and include any relevant examples from your own knowledge or experience. Write at least 250 words.";
const article2 = " if you take risks in your personal life,  you may discover new passions or hobbies that you never knew existed. However, taking risks also has its drawbacks. For example, if you take risks in your professional life and fail, you may lose your job or damage your reputation. Likewise, if you take risks in your personal life and fail, you may suffer from emotional or financial consequences.\\n\\nIn my opinion, the advantages of taking risks outweigh the disadvantages. Without taking risks, we would never be able to achieve great things or experience new opportunities. Of course, we should always weigh the potential benefits against the potential drawbacks before taking risks, and we should be prepared to accept the consequences of our actions. Overall, taking risks is an important aspect of life, and it can lead to personal and professional growth and fulfillment.";

const similarity = calculateSimilarity(article1, article2);
const Similarity = calSimilarity(article1, article2);

console.log(`article similarity：${similarity.toFixed(2)}`);
console.log(`jaccard similarity：${Similarity}`);

console.log((parseFloat(similarity.toFixed(2)) + parseFloat(Similarity)) / 2)

if (similarity.toFixed(2) > 0.4 && Similarity > 0.5) {
    console.log("high similarity")
} else {
    console.log("normal")
}
