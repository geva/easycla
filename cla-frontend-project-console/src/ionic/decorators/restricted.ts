// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

import { tap } from "rxjs/operators";

export function Restricted(restrictions: any) {
  return function (target: Function) {
    target.prototype.ionViewCanEnter = function () {
      return true;
    };
  };
}
