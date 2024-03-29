#include "console_logger.hpp"

int main(void)
{
	auto logger = logtard::ConsoleLogger();
	LOG_INFO(logger, "Hello, world!");
}
