import 'package:flutter/material.dart';
import 'package:flutter/src/foundation/key.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:photos/ente_theme_data.dart';

class MenuItemWidget extends StatelessWidget {
  final String text;
  final String? subText;
  final TextStyle? textStyle;
  final Color? leadingIconColor;
  final bool isBigger;
  const MenuItemWidget({
    required this.text,
    this.subText,
    this.textStyle,
    this.leadingIconColor,
    this.isBigger = false,
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final enteTheme = Theme.of(context).colorScheme.enteTheme;
    return Container(
      height: isBigger ? 48 : 44,
      width: double.infinity,
      padding: const EdgeInsets.symmetric(horizontal: 12),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Icon(
                Icons.add_outlined,
                size: 20,
                color: leadingIconColor ?? enteTheme.colorScheme.strokeBase,
              ),
              const SizedBox(width: 12),
              Text(
                text,
                style: textStyle ?? enteTheme.textTheme.bodyBold,
              ),
              subText != null
                  ? Row(
                      children: [
                        Padding(
                          padding: const EdgeInsets.symmetric(horizontal: 4),
                          child: Text(
                            '\u2022',
                            style: enteTheme.textTheme.small.copyWith(
                              color: enteTheme.colorScheme.textMuted,
                            ),
                          ),
                        ),
                        Text(
                          subText!,
                          style: enteTheme.textTheme.small
                              .copyWith(color: enteTheme.colorScheme.textMuted),
                        ),
                      ],
                    )
                  : const SizedBox.shrink(),
            ],
          ),
          Icon(Icons.chevron_right),
        ],
      ),
    );
  }
}
