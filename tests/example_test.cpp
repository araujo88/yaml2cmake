#include <gtest/gtest.h>

int add(int a, int b) {
    return a + b;
}

TEST(AdditionTest, HandlesPositiveInput) {
    EXPECT_EQ(3, add(1, 2));
}

TEST(AdditionTest, HandlesNegativeInput) {
    EXPECT_EQ(-1, add(-1, 0));
}
