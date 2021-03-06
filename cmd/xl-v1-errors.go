/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import "errors"

// errXLMaxDisks - returned for reached maximum of disks.
var errXLMaxDisks = errors.New("Number of disks are higher than supported maximum count '16'")

// errXLMinDisks - returned for minimum number of disks.
var errXLMinDisks = errors.New("Minimum '6' disks are required to enable erasure code")

// errXLNumDisks - returned for odd number of disks.
var errXLNumDisks = errors.New("Total number of disks should be multiples of '2'")

// errXLReadQuorum - did not meet read quorum.
var errXLReadQuorum = errors.New("Read failed. Insufficient number of disks online")

// errXLWriteQuorum - did not meet write quorum.
var errXLWriteQuorum = errors.New("Write failed. Insufficient number of disks online")
