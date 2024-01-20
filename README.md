## How to go: Advance go concepts following best practices


### Composability in Go (Isn't that all go offers )
In the provided code, the `hash` method and the `hash` function in the `HashReader` interface refer to different things.

1. **`hash` Method in `hashReader` Struct:**
   - The `hash` method is a method defined on the `hashReader` type. This method is specific to the `hashReader` type and is not part of the `HashReader` interface. It calculates the hash of the internal buffer (`buf`) and returns the result as a hex-encoded string. This method is not accessible through the `HashReader` interface.

   ```go
   func (h *hashReader) hash() string {
       return hex.EncodeToString(h.buf.Bytes())
   }
   ```

2. **`hash` Function in `HashReader` Interface:**
   - The `hash` function is part of the `HashReader` interface. The interface declares that any type implementing `HashReader` must have a method named `hash()` that takes no arguments and returns a string. This allows you to treat any type conforming to the `HashReader` interface uniformly, even if the underlying implementation details differ.

   ```go
   type HashReader interface {
       io.Reader
       hash() string
   }
   ```

   The `HashAndBroadCast` function uses this interface and calls the `hash()` method on the provided `HashReader`. This function can work with any type that implements the `HashReader` interface, allowing for flexibility and composability.

   ```go
   func HashAndBroadCast(r HashReader) error {
       hash := r.hash()
       // ... rest of the function ...
   }
   ```

In summary, the `hash` method is a specific implementation detail of the `hashReader` type, while the `hash` function is part of the `HashReader` interface, providing a common interface for different types. This enables polymorphism and allows the `HashAndBroadCast` function to work with any type adhering to the `HashReader` interface.
