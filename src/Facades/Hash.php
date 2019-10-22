<?php
declare(strict_types=1);

namespace Hyperf\Support\Facades;

use Hyperf\Support\Common\Facade;
use Hyperf\Extra\Contract\HashServiceInterface;

/**
 * Class Hash
 * @package Hyperf\Support\Facades
 * @method static string make(string $password, array $options = [])
 * @method static boolean check(string $password, string $hash)
 */
final class Hash extends Facade
{
    protected static function getFacadeAccessor()
    {
        return HashServiceInterface::class;
    }
}