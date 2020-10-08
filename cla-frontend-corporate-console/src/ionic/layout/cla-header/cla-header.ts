// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

import { Component, EventEmitter, Input, Output } from '@angular/core';
import { EnvConfig } from '../../services/cla.env.utils';

@Component({
  selector: 'cla-header',
  templateUrl: 'cla-header.html'
})
export class ClaHeader {
  @Input() title = '';
  @Input() hasShowMenu = true;
  @Output() onToggle: EventEmitter<any> = new EventEmitter<any>();
 
  hasEnabledLFXHeader = EnvConfig['lfx-header-enabled'];
  hasExpanded: boolean = true;
  
  onToggled() {
    this.hasExpanded = !this.hasExpanded;
    this.onToggle.emit(this.hasExpanded);
  }
}

