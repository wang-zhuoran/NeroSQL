package nerosql

import "math"

// LinearRegression predicts the output based on the input features using linear regression
func LinearRegression(features [][]float64, label []float64, optimizer string, learning_rate float64, epochs int) []float64 {
	numFeatures := len(features[0])
	numSamples := len(features)

	// Initialize weights with zeros
	weights := make([]float64, numFeatures)

	// Gradient descent to update weights
	alpha := learning_rate
	beta := 0.9 // used for momentum optimizer, 不想费劲让用户自定义了，就先这样吧。。。
	numIterations := epochs
	var prevGradient []float64
	for iter := 0; iter < numIterations; iter++ {
		for i := 0; i < numSamples; i++ {
			prediction := 0.0
			for j := 0; j < numFeatures; j++ {
				prediction += weights[j] * features[i][j]
			}
			error := label[i] - prediction
			gradient := make([]float64, numFeatures)
			for j := 0; j < numFeatures; j++ {
				gradient[j] = error * features[i][j]
				if optimizer == "sgd" {
					weights[j] += alpha * gradient[j]
				} else if optimizer == "momentum" {
					if iter == 0 {
						prevGradient = make([]float64, numFeatures)
					}
					prevGradient[j] = alpha*gradient[j] + beta*prevGradient[j]
					weights[j] += prevGradient[j]
				} else if optimizer == "adam" {
					beta1 := 0.9                                   // Exponential decay rate for the first moment estimates
					beta2 := 0.999                                 // Exponential decay rate for the second moment estimates
					epsilon := 1e-8                                // A small constant for numerical stability
					var m []float64 = make([]float64, numFeatures) // First moment vector
					var v []float64 = make([]float64, numFeatures) // Second moment vector
					t := iter + 1                                  // Bias correction term
					for j := 1; j < numFeatures; j++ {
						m[j-1] = beta1*m[j-1] + (1-beta1)*gradient[j]             // Update first moment estimate
						v[j-1] = beta2*v[j-1] + (1-beta2)*gradient[j]*gradient[j] // Update second moment estimate
						mHat := m[j-1] / (1 - math.Pow(beta1, float64(t)))        // Compute bias-corrected first moment estimate
						vHat := v[j-1] / (1 - math.Pow(beta2, float64(t)))        // Compute bias-corrected second moment estimate
						weights[j] += alpha * mHat / (math.Sqrt(vHat) + epsilon)  // Update weights
					}
				} else {
					panic("Invalid optimizer")
				}
			}
		}
	}

	return weights
}
