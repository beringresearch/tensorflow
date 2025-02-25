/* Copyright 2017 The TensorFlow Authors. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
==============================================================================*/

#ifndef TENSORFLOW_COMPILER_XLA_SERVICE_GPU_FFT_THUNK_H_
#define TENSORFLOW_COMPILER_XLA_SERVICE_GPU_FFT_THUNK_H_

#include "absl/container/flat_hash_map.h"
#include "absl/types/optional.h"
#include "tensorflow/compiler/xla/service/buffer_assignment.h"
#include "tensorflow/compiler/xla/service/gpu/buffer_allocations.h"
#include "tensorflow/compiler/xla/service/gpu/gpu_executable.h"
#include "tensorflow/compiler/xla/service/gpu/thunk.h"
#include "tensorflow/compiler/xla/service/hlo_instruction.h"
#include "tensorflow/compiler/xla/types.h"
#include "tensorflow/compiler/xla/xla_data.pb.h"
#include "tensorflow/core/lib/core/status.h"
#include "tensorflow/core/platform/stream_executor_no_cuda.h"

namespace xla {
namespace gpu {

// This class stores everything that StreamExecutor needs to launch an FFT.
// It is generated by IrEmitter.
//
// This is thread-compatible.
class FftThunk : public Thunk {
 public:
  // Constructs a thunk for launching an FFT on a stream.
  // Semantics of null hlo_instruction argument are as in Thunk.
  FftThunk(ThunkInfo thunk_info, FftType fft_type,
           absl::Span<const int64_t> fft_length,
           const BufferAllocation::Slice& input_buffer,
           const BufferAllocation::Slice& output_buffer,
           const Shape& input_shape, const Shape& output_shape);

  FftThunk(const FftThunk&) = delete;             // Cannot share fft_plan_
  FftThunk& operator=(const FftThunk&) = delete;  // Cannot share fft_plan_

  // Does the FFT for the thunk on "stream".
  Status ExecuteOnStream(const ExecuteParams& params) override;

 private:
  const se::fft::Type fft_type_;
  const std::vector<int64_t> fft_length_;

  float scale_factor_;

  // One plan per device ordinal.
  absl::Mutex mu_;
  struct FftPlan {
    absl::Mutex mu;
    std::unique_ptr<se::fft::Plan> plan;
  };
  absl::flat_hash_map<int, std::unique_ptr<FftPlan>> fft_plans_
      ABSL_GUARDED_BY(mu_);

  const BufferAllocation::Slice input_buffer_;
  const BufferAllocation::Slice output_buffer_;

  const Shape input_shape_;
  const Shape output_shape_;
};

}  // namespace gpu
}  // namespace xla

#endif  // TENSORFLOW_COMPILER_XLA_SERVICE_GPU_FFT_THUNK_H_
