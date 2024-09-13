// Copyright 2024 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

import classNames from "classnames/bind";
import React, { ReactNode } from "react";

import styles from "./pageSection.module.scss";
const cx = classNames.bind(styles);

type PageSectionProps = {
  heading?: string | ReactNode;
  children: ReactNode;
  className?: string;
  childrenClassname?: string;
};

export const PageSection: React.FC<PageSectionProps> = ({
  heading,
  children,
  className,
  childrenClassname,
}) => {
  const headingEl =
    heading && typeof heading === "string" ? (
      <h2 className={cx("page-section__heading")}>{heading}</h2>
    ) : (
      heading
    );

  return (
    <div className={cx("page-section", className)}>
      {headingEl}
      <div className={cx("page-section__content", childrenClassname)}>
        {children}
      </div>
    </div>
  );
};
