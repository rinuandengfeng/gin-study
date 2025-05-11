#!/usr/bin/env awk

{
  print $0
  if (match($0, /^total:/)) {
    sub(/%/, "", $NF);
    printf("测试覆盖率为 %s%(要求达到 %s%)\n", $NF, target)
    if (strtonum($NF) < target) {
      printf("测试覆盖率未达到预期: %d%, 请添加测试用例!\n", target)
      exit 1;
    }
  }
}
