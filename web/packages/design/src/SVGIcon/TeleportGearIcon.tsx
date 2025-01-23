/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import React from 'react';
import { useTheme } from 'styled-components';

import type { SVGIconProps } from './common';
import { SVGIcon } from './SVGIcon';

export function TeleportGearIcon({ size = 16 }: Omit<SVGIconProps, 'fill'>) {
  const theme = useTheme();
  let fill = theme.colors.brand;

  // All other themes will default to white.
  if (theme.type !== 'light') {
    fill = 'white';
  }

  return (
    <SVGIcon size={size} viewBox="0 0 164 164">
      <g id="logos" fill={fill}>
        <path
          fillRule="evenodd"
          clipRule="evenodd"
          d="M49.2716 82.0054C49.2716 63.9206 64.0081 49.2064 82.1206 49.2051C100.234 49.2064 114.972 63.9207 114.972 82.0054C114.972 100.09 100.236 114.803 82.1219 114.803C64.0082 114.803 49.2716 100.09 49.2716 82.0054ZM62.7955 82.0054C62.7955 92.6332 71.465 101.279 82.1219 101.279C92.7787 101.279 101.448 92.6332 101.448 82.0054C101.448 71.3762 92.7759 62.7295 82.1178 62.7295C71.4637 62.7295 62.7955 71.3762 62.7955 82.0054Z"
        />
        <path
          fillRule="evenodd"
          clipRule="evenodd"
          d="M144.017 91.5473L160.009 105.164C160.563 105.556 160.962 106.13 161.142 106.786C161.316 107.443 161.264 108.144 160.981 108.762C157.723 118.619 152.103 127.98 145.163 135.737C144.288 136.735 143.103 137.167 141.931 136.851L141.493 136.703L121.351 129.814C116.929 133.375 111.989 136.31 106.714 138.538L104.718 139.343L100.721 159.841C100.584 160.511 100.244 161.122 99.7478 161.599C99.2522 162.069 98.6236 162.377 97.9472 162.487C92.7267 163.44 87.5059 164 82.1278 164C76.7439 164 71.5234 163.44 66.3028 162.487C65.6982 162.39 65.131 162.14 64.6617 161.747C64.1917 161.354 63.8377 160.839 63.6374 160.266L63.5278 159.841L59.5376 139.343C54.1079 137.302 48.9986 134.489 44.3679 130.999L42.8988 129.814L22.7546 136.703C21.456 137.263 20.0691 136.858 19.0848 135.737C12.1491 127.98 6.52575 118.619 3.26619 108.762C3.02161 108.215 2.95011 107.604 3.06081 107.011C3.17204 106.419 3.4606 105.878 3.88812 105.453L4.24339 105.164L20.2264 91.5473C19.69 88.3995 19.4434 85.2124 19.4877 82.0193C19.4877 79.4636 19.5907 76.843 19.9256 74.3451L20.2264 72.4974L4.23622 58.8942C3.6812 58.5013 3.2809 57.9221 3.10279 57.2655C2.92418 56.6087 2.97963 55.9132 3.25906 55.2951C6.51862 45.4452 12.142 36.0716 19.0782 28.3008C19.4561 27.831 19.9602 27.4771 20.5333 27.2776C21.1059 27.0839 21.7217 27.0521 22.3103 27.1938L22.7546 27.3353L42.8988 34.2173C47.3801 30.5671 52.3769 27.6058 57.7267 25.4168L59.5376 24.6959L63.5278 4.21025C63.6644 3.5344 64.0028 2.92264 64.4993 2.44651C64.9958 1.9765 65.6249 1.66731 66.3028 1.56424C76.7451 -0.521413 87.4969 -0.521413 97.94 1.56424L97.9472 1.57087C98.5503 1.66097 99.1167 1.91821 99.5855 2.31113C100.054 2.70375 100.406 3.21245 100.605 3.7919L100.721 4.21688L104.712 24.7021C110.05 26.7498 115.082 29.5177 119.67 32.9297L121.344 34.2237L141.493 27.3483C142.794 26.7816 144.178 27.1875 145.157 28.3008C152.084 36.0717 157.717 45.4452 160.981 55.2952C161.225 55.8425 161.303 56.454 161.187 57.0464C161.078 57.6325 160.788 58.1796 160.363 58.598L160.009 58.8942L144.017 72.491C144.59 75.5297 144.757 78.8192 144.757 82.0193C144.757 84.5945 144.648 87.2016 144.313 89.6995L144.017 91.5473ZM34.0425 82.0063C34.0425 108.505 55.5673 129.988 82.1214 129.988C108.675 129.988 130.201 108.505 130.201 82.0063C130.201 55.5091 108.677 34.0201 82.1178 34.018C55.552 34.0201 34.0425 55.5091 34.0425 82.0063Z"
        />
      </g>
    </SVGIcon>
  );
}
